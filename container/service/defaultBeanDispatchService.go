package service

import (
	"github.com/MajorShebka/di-container/container/service/model"
	"reflect"
)

type DefaultBeanDispatchService struct {
}

func (service DefaultBeanDispatchService) Dispatch(tag string, nameToBean map[string]interface{}, t reflect.Type, kind reflect.Kind) model.Bean {
	beanName := tag
	if kind == reflect.Interface {
		beanName += "struct"
	} else {
		beanName += kind.String()
	}
	return model.Bean{nameToBean[beanName], nil, t}
}

func (service DefaultBeanDispatchService) GetType() string {
	return "default"
}
