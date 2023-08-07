package repository

import (
	"context"
	"fmt"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
)

type projectRepository struct {
	database   *qmgo.Database
	collection string
}

func NewProjectRepository(db *qmgo.Database) domain.ProjectRepository {
	return &projectRepository{
		database:   db,
		collection: domain.CollectionPorject,
	}
}

func (p *projectRepository) Create(c context.Context, project domain.Project) error {
	collection := p.database.Collection(p.collection)
	_, err := collection.InsertOne(c, project)
	if err != nil {
		return fmt.Errorf("failed to save the project,project:%v; %v", project, err)
	}
	bootstrap.Logger.Debugf("insert prject success value: %v", project)
	return nil
}

func (p *projectRepository) QueryByName(c context.Context, name string) ([]domain.Project, error) {
	collection := p.database.Collection(p.collection)
	batch := []domain.Project{}
	if err := collection.Find(c, bson.M{"projectName": name}).Sort("version").All(&batch); err != nil {
		return nil, fmt.Errorf("Query project by Name:%s err %v", name, err)
	}
	return batch, nil
}
