package adapt

/*
把*proto.Message类型的返回值转换成 map[string]interface{}
方便快速的载入 gota (github.com/go-gota/gota/dataframe), gota是GoLang下的计算工具，与Python下的pandas类似
转换后的key，是使用的protobuf中的`json`tag的定义，首字母小写的那个
你也可以使用adapt.Field('oldkey','newkey')替换掉旧的key，因为有些字段定义太长，在gota下会被折叠，或者与你自己的数据库字段定义不一致
useage:
adapt.PbParser(
		adapt.Field("time", "timeChange"),
		adapt.Field("key", "keyNew"),
	)
一旦使用了 adapt.Field('oldkey','newkey')，就需要把所有需要的key list都加入，没有加入的key会被丢弃，可以使用这个来过滤不需要的key
如果只是过滤key，不改变key的值，newkey可以传空字符串
*/
import (
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
)

//===
// field set
type field struct {
	key   string
	alias string
}

func Field(key string, alias string) field {
	if alias == "" {
		alias = key
	}
	return field{
		key:   key,
		alias: alias,
	}
}

func (p *PBMessageParser) setFields(a field) {
	p.fields[a.key] = a.alias
}

//===
//todo custom change
var customMap = map[string]interface{}{
	"qotcommon.Security": func(pm *proto.Message) {},
}

//===

type PBMessageParser struct {
	fields map[string]string
}

func PbParser(fields ...field) *PBMessageParser {
	pm := &PBMessageParser{
		fields: make(map[string]string),
	}
	for _, f := range fields {
		pm.setFields(f)
	}
	return pm
}

//Todo fix type
func (p *PBMessageParser) ArrMap(pms []proto.Message) []map[string]interface{} {
	mp := make([]map[string]interface{}, 0)
	for _, v := range pms {
		mp = append(mp, p.Map(v))
	}
	return mp
}

func (p *PBMessageParser) Map(pm proto.Message) map[string]interface{} {
	return p.parseStruct(pm)
}

func (p *PBMessageParser) parseValue(fvalue reflect.Value) interface{} {
	switch fvalue.Kind() {
	case reflect.Slice, reflect.Array:
		m := make([]interface{}, 0)
		for i := 0; i < fvalue.Len(); i++ {
			aValue := fvalue.Index(i)
			if aValue.Kind() == reflect.Ptr {
				aValue = aValue.Elem()
			}
			//递归
			m = append(m, p.parseValue(aValue))
		}
		return m
	case reflect.Struct:
		//递归，新建一个对象，因为要丢掉Field设置
		return PbParser().parseStruct(fvalue.Interface())
	case reflect.Map:
		return fvalue.Interface()
	case reflect.Chan:
		return fvalue.Interface()
	case reflect.Int64:
		return fvalue.Int()
	case reflect.Int8:
		return int8(fvalue.Int())
	case reflect.Int16:
		return int16(fvalue.Int())
	case reflect.Int32:
		return int32(fvalue.Int())
	case reflect.Int:
		return int(fvalue.Int())
	case reflect.Uint64:
		return fvalue.Uint()
	case reflect.Uint8:
		return uint8(fvalue.Uint())
	case reflect.Uint16:
		return uint16(fvalue.Uint())
	case reflect.Uint32:
		return uint32(fvalue.Uint())
	case reflect.Uint:
		return uint(fvalue.Uint())
	case reflect.Float64:
		return fvalue.Float()
	case reflect.Float32:
		return float32(fvalue.Float())
	case reflect.String:
		return fvalue.String()
	case reflect.Bool:
		return fvalue.Bool()
	case reflect.Complex64, reflect.Complex128:
		return fvalue.Complex()
	case reflect.Interface:
		return fvalue.Interface()
	default:
		return nil
	}
}

func (p *PBMessageParser) parseStruct(pm interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	//todo fix types

	v := reflect.ValueOf(pm)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return res
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// only accept struct param
	if v.Kind() != reflect.Struct {
		return res
	}
	//todo Struct 自定义转换

	tag := "json"
	t := v.Type()
	if in(t.Name(), []string{"qotcommon.Security"}) {
		//todo fix some pb
	}
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		// read tag
		val, ok := fieldType.Tag.Lookup(tag)
		if !ok { //没有json tag 跳过
			continue
		}
		opts := strings.Split(val, ",")
		if len(opts) < 2 { //json tag字段少于2 跳过
			continue
		}
		field_key := opts[0]
		field_option := opts[1:]

		fieldValue := v.Field(i)
		if in("omitempty", field_option) && fieldValue.IsZero() {
			continue
		}
		// ignore nil pointer in field
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}
		pv := p.parseValue(fieldValue)
		if nil != pv {
			//todo field_change
			if len(p.fields) > 0 {
				if alias_key, ok := p.fields[field_key]; ok {
					field_key = alias_key
				} else {
					continue
				}
			}
			res[field_key] = pv
		}
	}
	return res
}

func in(e string, arr []string) bool {
	for _, v := range arr {
		if v == e {
			return true
		}
	}
	return false
}
