package service

import (
	"github.com/MajorShebka/di-container/container/util"
	"reflect"
)

type InjectService interface {
	Inject(interface{}, reflect.StructField, reflect.Value, map[string]interface{})
	IsSuitable(reflect.StructTag) bool
}

func GetInjectService(tag reflect.StructTag) InjectService {
	if tag.Get(util.METHOD_TAG) != "" {
		return MethodInjectService{}
	}
	return DefaultInjectService{}
}
