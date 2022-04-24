package pkg

import (
	"fmt"
	"reflect"
)

// ConvertStruct 将source强转为typ类型（返回值为新建typ类型对象）
func ConvertStruct(source any, typ *any) (s any, err error) {
	defer func() {
		if _err := recover(); _err != nil {
			fmt.Printf("Runtime panic caught: %v\n", _err)
			err = fmt.Errorf("Runtime panic caught: %v\n", _err)
		}
	}()

	t := reflect.TypeOf(*typ)
	v := reflect.ValueOf(source)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	newStruct := reflect.New(t)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Name
		value := v.FieldByName(field)
		tmp := newStruct.Elem().Field(i)
		tmp.Set(value)
	}

	s = newStruct.Interface()
	return
}
