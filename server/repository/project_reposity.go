package repository

import (
	"context"

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
	return err
}
