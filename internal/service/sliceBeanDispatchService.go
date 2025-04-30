package service

import (
	"DI-container/internal/service/model"
	"DI-container/internal/util"
	"reflect"
	"strings"
)

type SliceBeanDispatchService struct {
}

func (service *SliceBeanDispatchService) Dispatch(tag string, nameToBean map[string]interface{}, t reflect.Type, kind reflect.Kind) model.Bean {
	slice := make([]interface{}, 0)
	beans := strings.Split(tag, util.COMMA)
	for _, bean := range beans {
		existedBean := nameToBean[bean+"struct"]
		if existedBean != nil {
			slice = append(slice, existedBean)
		}
	}
	return model.Bean{nil, slice, t}
}

func (service *SliceBeanDispatchService) GetType() string {
	return "slice"
}
