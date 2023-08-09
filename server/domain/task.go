package domain

import (
	"context"
)

const (
	CollectionTask = "task"
	TaskStarting   = 1
	TaskSuccess    = 0
	TaskFaild      = -1
)

type Task struct {
	Name    string `bson:"name"`
	Status  string `bson:"status"`
	Version string `bson:"version"`
}

type TaskRepository interface {
	Create(c context.Context, task Task) error
	Update(c context.Context, task Task) (Task, error)
	Query(c context.Context, id string) (Task, error)
}
