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

func NewTaskRepository(db *qmgo.Database) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: domain.CollectionTask,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	return err
}

func (tr *taskRepository) QueryUpdateById(c context.Context, task *domain.Task, id string) (domain.Task, error) {
	return domain.Task{}, nil
}
