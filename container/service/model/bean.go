package model

import (
	"github.com/MajorShebka/di-container/container/util"
	"reflect"
)

type Bean struct {
	Single interface{}
	Multi  []interface{}
	T      reflect.Type
}

func (b Bean) GetValue() reflect.Value {
	if b.Single != nil {
		return reflect.ValueOf(b.Single)
	}
	if b.Multi != nil {
		slice := reflect.MakeSlice(reflect.SliceOf(b.T), 0, len(b.Multi))
		for _, v := range b.Multi {
			slice = reflect.Append(slice, util.GetValue(v))
		}
		return slice
	}
	return reflect.ValueOf(nil)
}
