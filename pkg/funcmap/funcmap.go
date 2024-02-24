package funcmap

import (
	"html/template"
	"reflect"

	"github.com/Masterminds/sprig"
)

func GetFuncMap() template.FuncMap {
	funcMap := sprig.FuncMap()

	f := template.FuncMap{
		"hasField": HasField,
	}
	for k, v := range f {
		funcMap[k] = v
	}
	return funcMap
}

func HasField(v any, field string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(field).IsValid()
}
