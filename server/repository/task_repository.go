package repository

import (
	"context"

	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
)

type taskRepository struct {
	database   *qmgo.Database
	collection string
}

func NewTaskRepository(db *qmgo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	return nil
}

func (tr *taskRepository) QueryUpdateById(c context.Context, task *domain.Task, id string) (domain.Task, error) {
	return domain.Task{}, nil
}
