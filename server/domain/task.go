package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "task"
	TaskStarting   = 1
	TaskSuccess    = 0
	TaskFaild      = -1
)

type Task struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Status string             `bson:"status"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	Update(c context.Context, task *Task) (Task, error)
}
