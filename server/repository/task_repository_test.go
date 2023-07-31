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
	objId, err := primitive.ObjectIDFromHex("64c7a43dea4c2e21fbbbfe63")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	task := domain.Task{
		ID:     objId,
		Status: "-2",
	}
	_, err = tr.Update(context.Background(), &task)
	bootstrap.Logger.Error(err)
}

func TestCreate(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)
	task := domain.Task{
		Name:   "one",
		Status: "1",
	}
	err := tr.Create(context.Background(), &task)
	bootstrap.Logger.Error(err)
}

func TestQuery(t *testing.T) {
	db := app.Mongo.Database("groot")
	tr := repository.NewTaskRepository(db)

	rs, err := tr.Query(context.Background(), "64c7a43dea4c2e21fbbbfe63")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	bootstrap.Logger.Info(rs)
}
