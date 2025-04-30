package service

import (
	"github.com/MajorShebka/di-container/container/util"
	"reflect"
	"strings"
)

type DefaultInjectService struct {
}

func (service DefaultInjectService) Inject(obj interface{}, field reflect.StructField, fieldValue reflect.Value, nameToBean map[string]interface{}) {
	dispatchService := GetBeanDispatchService(field)
	tag := field.Tag.Get(util.INJECT_TAG)
	bean := dispatchService.Dispatch(tag, nameToBean, field.Type, fieldValue.Kind()).GetValue()
	if bean == reflect.ValueOf(nil) {
		return
	}
	if fieldValue.CanSet() {
		fieldValue.Set(bean)
	} else {
		method := reflect.ValueOf(obj).MethodByName(service.createSetterName(field))
		if method.IsValid() {
			method.Call([]reflect.Value{bean})
		}
	}
}

func (service DefaultInjectService) IsSuitable(tag reflect.StructTag) bool {
	return true
}

func (service DefaultInjectService) createSetterName(field reflect.StructField) string {
	fieldName := field.Name
	return "Set" + strings.ToUpper(fieldName[0:1]) + fieldName[1:]
}
