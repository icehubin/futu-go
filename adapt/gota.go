package adapt

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
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		return fvalue.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		return fvalue.Uint()
	case reflect.Float32, reflect.Float64:
		return fvalue.Float()
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
