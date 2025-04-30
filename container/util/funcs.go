package util

import "reflect"

func GetValue(obj interface{}) reflect.Value {
	value := reflect.ValueOf(obj)
	if value.Kind() != reflect.Ptr {
		return value
	}
	return value.Elem()
}

func TypeOf(obj interface{}) reflect.Type {
	value := reflect.TypeOf(obj)
	if value.Kind() != reflect.Ptr {
		return value
	}
	return value.Elem()
}

func SliceOf(t reflect.Type) reflect.Type {
	return reflect.SliceOf(t).Elem()
}

func CreateKey(t reflect.Type) string {
	return t.Name() + t.Kind().String()
}
