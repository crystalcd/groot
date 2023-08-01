package repository_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/repository"
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
	testobj := TestBuildStruct{
		Task:    task,
		TaskPtr: &task,
		Ints:    []int{1, 2, 3},
		Strings: []string{"hello", "world"},
		i:       123,
	}
	update := repository.BuildUpdate(testobj)
	bootstrap.Logger.Info(update)
}

type TestBuildStruct struct {
	Task    domain.Task
	TaskPtr *domain.Task
	Ints    []int
	Strings []string
	i       int
}

type Inner struct {
}

func TestRefeact(t *testing.T) {
	task := domain.Task{
		Status: "1",
	}

	type1 := reflect.TypeOf(task)
	typeptr := reflect.TypeOf(&task)
	tp := type1.Kind()
	prtkind := typeptr.Kind()
	value := reflect.ValueOf(task)
	valueptr := reflect.ValueOf(&task)
	bootstrap.Logger.Debug(type1, tp)
	bootstrap.Logger.Debug(typeptr, prtkind)
	bootstrap.Logger.Debug(value)
	bootstrap.Logger.Debug(valueptr)

	for i := 0; i < type1.NumField(); i++ {
		bootstrap.Logger.Debug("FieldName:", type1.Field(i).Name, "FieldType:", type1.Field(i).Type, "FieldValue:", value.Field(i))
	}

	bootstrap.Logger.Debugln("-----------------------------------------------------")

	for i := 0; i < value.NumField(); i++ {
		bootstrap.Logger.Debug("FieldName:", type1.Field(i).Name, "FieldType:", type1.Field(i).Type, "FieldValue:", value.Field(i))
	}

	zero := reflect.Zero(type1)
	bootstrap.Logger.Debug(zero)
}
