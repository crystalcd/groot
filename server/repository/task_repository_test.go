package repository_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var app bootstrap.Application

func TestMain(m *testing.M) {
	app = bootstrap.App()
	m.Run()
}

func TestUpdate(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)
	objId, err := primitive.ObjectIDFromHex("64c85f3a3f3ccb00c2f93fac")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	task := domain.Task{
		ID:     objId,
		Status: "0",
	}
	_, err = tr.Update(context.Background(), &task)
	bootstrap.Logger.Error(err)
}

func TestCreate(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)
	task := domain.Task{
		ID:     primitive.NewObjectID(),
		Name:   "one",
		Status: "1",
	}
	err := tr.Create(context.Background(), &task)
	bootstrap.Logger.Error(err)
}

func TestQuery(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)

	rs, err := tr.Query(context.Background(), "64c85f3a3f3ccb00c2f93fac")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	bootstrap.Logger.Info(rs)
}

func TestBuild(t *testing.T) {
	task := domain.Task{
		Name: "AA",
	}
	update := BuildUpdate(task)
	bootstrap.Logger.Info(update)
}

func BuildUpdate(i interface{}) bson.M {
	return buildUpdate(i, true)
}

func buildUpdate(i interface{}, ignoreID bool) bson.M {
	update := bson.M{}

	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		if isZero(fieldVal) {
			continue
		}

		fieldName := t.Field(i).Tag.Get("bson")

		// 忽略_id字段
		if ignoreID && fieldName == "_id" {
			continue
		}

		if fieldVal.Kind() == reflect.Struct {
			update[fieldName] = buildUpdate(fieldVal.Interface(), ignoreID)
		} else {
			update[fieldName] = fieldVal.Interface()
		}
	}

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
