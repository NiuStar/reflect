package reflect

import (
	"reflect"
	//"fmt"
)

func GetReflectType(value interface{}) reflect.Type {
	v := reflect.TypeOf(value)

	for ;reflect.Ptr == v.Kind(); {

		//fmt.Println("v = v.Elem():",v.String())
		v = v.Elem()
	}
	return v
}

func GetReflectValue(value interface{}) reflect.Value {
	v := reflect.ValueOf(value)

	for ;reflect.Ptr == v.Kind()||reflect.Interface == v.Kind(); {
		v = v.Elem()
	}
	return v
}
