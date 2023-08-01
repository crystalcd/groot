package repository

import (
	"reflect"

	"github.com/crystal/groot/bootstrap"
	"go.mongodb.org/mongo-driver/bson"
)

func BuildUpdate(i interface{}) bson.M {
	return buildUpdate(i, true)
}

func buildUpdate(i interface{}, ignoreID bool) bson.M {
	update := bson.M{}
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	t.Kind()
	v.Kind()
	bootstrap.Logger.Debug(t, v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = reflect.Indirect(reflect.ValueOf(i)) // 处理指针
	}
	bootstrap.Logger.Debug(t, v)

	for i := 0; i < t.NumField(); i++ {
		fieldVal := v.Field(i)
		bootstrap.Logger.Debug("FieldValue:", fieldVal)
		if !fieldVal.CanInterface() {
			continue
		}

		fieldName := t.Field(i).Tag.Get("bson")
		if ignoreID && fieldName == "_id" {
			continue
		}

		if isZero(fieldVal) {
			continue
		}
		fieldType := t.Field(i)
		if fieldType.Type.Kind() == reflect.Ptr {
			actualType := fieldType.Type.Elem()
			if actualType.Kind() == reflect.Struct {
				update[fieldName] = buildUpdate(fieldVal.Interface(), ignoreID)
			}
		} else if fieldType.Type.Kind() == reflect.Struct {
			update[fieldName] = buildUpdate(fieldVal.Interface(), ignoreID)
		} else {
			update[fieldName] = fieldVal.Interface()
		}
	}

	bootstrap.Logger.Debugf("repository update json:%+v", update)
	return update
}

func isZero(v reflect.Value) bool {
	bootstrap.Logger.Debug("value kind: ", v.Kind(), v)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Struct:
		z := reflect.Zero(v.Type()).Interface()
		return reflect.DeepEqual(v.Interface(), z)
	}
	return false
}
