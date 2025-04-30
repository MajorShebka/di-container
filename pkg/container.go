package pkg

import (
	"DI-container/internal/service"
	"DI-container/internal/util"
	"reflect"
)

type Container struct {
	nameToBean map[string]interface{}
	beans      []interface{}
}

func NewContainer() *Container {
	return &Container{make(map[string]interface{}), make([]interface{}, 0)}
}

func (c *Container) Put(objects ...interface{}) {
	for i := range objects {
		c.beans = append(c.beans, objects[i])
		objType := util.TypeOf(objects[i])
		c.nameToBean[objType.Name()+reflect.ValueOf(objects[i]).Kind().String()] = objects[i]
	}
}

func (c *Container) Init() {
	for _, obj := range c.beans {
		objType := util.TypeOf(obj)
		c.inject(objType, obj)
	}
}

func (c *Container) inject(objType reflect.Type, obj interface{}) {
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if tag := field.Tag.Get(util.INJECT_TAG); tag != "" {
			fieldValue := util.GetValue(obj).Field(i)
			injectService := service.GetInjectService(field.Tag)
			injectService.Inject(obj, field, fieldValue, c.nameToBean)
		}
	}
}

func (c *Container) Get(bean interface{}) interface{} {
	beanType := util.TypeOf(bean)
	return c.nameToBean[beanType.Name()+reflect.ValueOf(beanType).Kind().String()]
}
