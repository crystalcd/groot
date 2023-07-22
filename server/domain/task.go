package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "task"
)

type Task struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Status string             `bson:"status"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	QueryUpdateById(c context.Context, task *Task, id string) (Task, error)
}
