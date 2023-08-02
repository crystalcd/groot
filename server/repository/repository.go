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
	update := make(bson.M)

	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		if v.IsNil() {
			return update
		}
		v = v.Elem()
	}

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)

		fieldName := v.Type().Field(i).Name
		fieldType := v.Type().Field(i).Type
		// 处理非导出字段
		if !fieldVal.CanInterface() {
			continue
		}
		bootstrap.Logger.Info(fieldVal.Kind())
		if fieldVal.Kind() == reflect.Ptr {
			if fieldVal.IsNil() {
				continue
			}
			fieldVal = fieldVal.Elem()
			fieldType = fieldVal.Type()
		}
		if IsZero(fieldVal) {
			continue
		}

		// 如果字段是结构体,递归处理
		if fieldType.Kind() == reflect.Struct {
			update[fieldName] = buildUpdate(fieldVal.Interface(), ignoreID)
			continue
		}

		update[fieldName] = fieldVal.Interface()
	}

	return update
}

func IsZero(v reflect.Value) bool {
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
