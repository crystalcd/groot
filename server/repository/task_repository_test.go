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

func TestIsZero(t *testing.T) {
	task := domain.Task{}
	cases := []struct {
		val  interface{}
		zero bool
	}{
		{0, true},
		{"", true},
		{nil, true},
		{[]int{}, true},
		{map[string]int{}, true},
		{struct{}{}, true},
		{1, false},
		{"foo", false},
		{map[string]int{"a": 1}, false},
		{(*int)(nil), true},
		{&task, false},
	}

	for _, c := range cases {
		if repository.IsZero(reflect.ValueOf(c.val)) != c.zero {
			t.Errorf("isZero %v failed", c.val)
		}
	}
}

type User struct {
	Name string
	Age  int
}

type Order struct {
	ID      int
	User    User
	UserPtr *User
	Items   []string
	i       int
}

func TestBuildUpdate1(t *testing.T) {
	o := Order{
		ID: 1,
		User: User{
			Name: "John",
			Age:  30,
		},

		Items: []string{"item1", "item2"},
	}

	// expected := map[string]interface{}{
	// 	"ID": 1,
	// 	"User": map[string]interface{}{
	// 		"Name": "John",
	// 		"Age":  30,
	// 	},
	// 	"Items": []string{"item1", "item2"},
	// }

	expected1 := bson.M{
		"ID": 1,
		"User": bson.M{
			"Name": "John",
			"Age":  30,
		},
		"Items": []string{"item1", "item2"},
	}

	actual := repository.BuildUpdate(o)
	// actual := buildUpdate(o)
	bootstrap.Logger.Info(actual)
	if !reflect.DeepEqual(expected1, actual) {
		t.Errorf("expected: %v, got: %v", expected1, actual)
	}
}

func buildUpdate(i interface{}) map[string]interface{} {

	result := make(map[string]interface{})

	v := reflect.ValueOf(i)

	// 遍历结构体字段
	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)

		// 处理非导出字段
		if !fieldVal.CanInterface() || repository.IsZero(v) {
			continue
		}

		fieldName := v.Type().Field(i).Name
		fieldType := v.Type().Field(i).Type

		// 如果字段是结构体,递归处理
		if fieldType.Kind() == reflect.Struct {
			result[fieldName] = buildUpdate(fieldVal.Interface())
			continue
		}

		result[fieldName] = fieldVal.Interface()
	}

	return result
}
