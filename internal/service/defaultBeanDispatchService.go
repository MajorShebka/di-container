package service

import (
	"DI-container/internal/service/model"
	"reflect"
)

type DefaultBeanDispatchService struct {
}

func (service DefaultBeanDispatchService) Dispatch(tag string, nameToBean map[string]interface{}, t reflect.Type, kind reflect.Kind) model.Bean {
	return model.Bean{nameToBean[tag+kind.String()], nil, t}
}

func (service DefaultBeanDispatchService) GetType() string {
	return "default"
}
