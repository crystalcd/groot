package repository

import (
	"reflect"

	"github.com/crystal/groot/bootstrap"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildUpdate(i interface{}) bson.M {
	return buildUpdate(reflect.Indirect(reflect.ValueOf(i)), true)
}

func buildUpdate(i interface{}, ignoreID bool) bson.M {

	update := bson.M{}

	v := reflect.Indirect(reflect.ValueOf(i)) // 处理指针

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)

		if !fieldVal.CanInterface() {
			continue
		}

		if isZero(fieldVal) {
			continue
		}

		fieldName := t.Field(i).Tag.Get("bson")

		if ignoreID && fieldName == "_id" {
			continue
		}

		if fieldVal.Kind() == reflect.Struct {
			update[fieldName] = buildUpdate(fieldVal.Interface(), ignoreID)
		} else {
			update[fieldName] = fieldVal.Interface()
		}
	}
	bootstrap.Logger.Debugf("repository update json:%+v", update)
	return update
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		// Others like int, float64 etc are default to their zero value
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}
