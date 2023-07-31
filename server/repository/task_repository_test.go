package repository_test

import (
	"context"
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
	task := domain.Task{
		ID:   primitive.ObjectID{},
		Name: "one",
	}
	_, err := tr.Update(context.Background(), &task)
	bootstrap.Logger.Error(err)
}

func TestCreate(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)
	task := domain.Task{
		ID:   primitive.NewObjectID(),
		Name: "one",
	}
	err := tr.Create(context.Background(), &task)
	bootstrap.Logger.Error(err)
}
