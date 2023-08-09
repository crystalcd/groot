package repository

import (
	"context"
	"fmt"

	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (tr *taskRepository) Create(c context.Context, task domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	if _, err := collection.InsertOne(c, task); err != nil {
		return fmt.Errorf("insert task taskId:%v err %v", task, err)
	}
	return nil
}

func (tr *taskRepository) Update(c context.Context, task domain.Task) (domain.Task, error) {

	return domain.Task{}, nil
}

func (tr *taskRepository) Query(c context.Context, id string) (domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	rs := domain.Task{}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rs, err
	}
	err = collection.Find(c, bson.M{"_id": objId}).One(&rs)
	return rs, err
}
