package repository

import (
	"context"
	"fmt"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
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
		return fmt.Errorf("failed to save the project,projectId:%s; %v", project.ProjectId, err)
	}
	bootstrap.Logger.Debugf("insert prject success value: %v", project)
	return nil
}
