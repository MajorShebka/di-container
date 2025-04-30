package service

import (
	"github.com/MajorShebka/go-di-container/internal/util"
	"reflect"
)

type MethodInjectService struct {
}

func (service MethodInjectService) Inject(obj interface{}, field reflect.StructField, fieldValue reflect.Value, nameToBean map[string]interface{}) {
	dispatchService := GetBeanDispatchService(field)
	bean := dispatchService.Dispatch(field.Tag.Get(util.INJECT_TAG), nameToBean, nil, 0)
	method := reflect.ValueOf(obj).MethodByName(field.Tag.Get(util.METHOD_TAG))
	if method.IsValid() {
		method.Call([]reflect.Value{bean.GetValue()})
	}
}

func (service MethodInjectService) IsSuitable(tag reflect.StructTag) bool {
	return true
}
