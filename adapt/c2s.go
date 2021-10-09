package adapt

import (
	"fmt"
	"reflect"
	"strings"
)

type Message map[string]interface{}

func protoAppend(pm interface{}, c Message) {
	fvalue := reflect.ValueOf(pm)
	if fvalue.Kind() == reflect.Ptr {
		fvalue = fvalue.Elem()
	}
	if fvalue.Kind() != reflect.Slice || !fvalue.CanSet() {
		return
	}
	st := fvalue.Type().Elem()
	for st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	nv := reflect.New(st)
	valueFill(nv, c)
	resArr := reflect.Append(fvalue, nv)
	if fvalue.CanSet() {
		fvalue.Set(resArr)
	}
}

func protoFill(pm interface{}, c Message) {
	fvalue := reflect.ValueOf(pm)
	valueFill(fvalue, c)
}

func valueFill(fvalue reflect.Value, c Message) {
	if fvalue.Kind() == reflect.Ptr {
		fvalue = fvalue.Elem()
	}
	fmt.Println(fvalue.Kind())
	// only accept struct param
	if fvalue.Kind() != reflect.Struct {
		return
	}
	tag := "json"
	t := fvalue.Type()
	// fmt.Println("type:", t)
	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := fvalue.Field(i)
		// read tag
		val, ok := fieldType.Tag.Lookup(tag)
		if !ok { //没有json tag 跳过
			continue
		}
		opts := strings.Split(val, ",")
		if len(opts) < 2 { //json tag字段少于2 跳过
			continue
		}
		// fmt.Println(fieldType, fieldValue.Type(), fieldValue.Kind())
		if fieldValue.Kind() == reflect.Ptr || fieldValue.Kind() == reflect.Slice {
			field_key := opts[0]
			if val2, ok2 := c[field_key]; ok2 {
				cvalue := reflect.ValueOf(val2)
				if cvalue.Type().String() == fieldValue.Type().String() {
					fieldValue.Set(cvalue)
				}
			}
		}
	}
}
