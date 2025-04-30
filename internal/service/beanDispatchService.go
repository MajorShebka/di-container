package service

import (
	"DI-container/internal/service/model"
	"reflect"
)

type BeanDispatchService interface {
	Dispatch(tag string, nameToBean map[string]interface{}, t reflect.Type, kind reflect.Kind) model.Bean
	GetType() string
}

func GetBeanDispatchService(field reflect.StructField) BeanDispatchService {
	switch field.Type.Kind() {
	case reflect.Slice:
		return &SliceBeanDispatchService{}
	default:
		return DefaultBeanDispatchService{}
	}
}
