package repository

import (
	"context"

	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
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

func (tr *taskRepository) Update(c context.Context, task *domain.Task) (domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	rs := domain.Task{}
	collection.UpdateId(c, task.ID.String(), task)
	return rs, nil
}

func (tr *taskRepository) QueryById(c context.Context, id string) (domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	rs := domain.Task{}
	err := collection.Find(c, bson.M{"id": id}).One(&rs)
	return rs, err
}
