package repository

import (
	"context"

	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
)

type projectReposity struct {
	database   *qmgo.Database
	collection string
}

func NewProjectReposity(db *qmgo.Database) domain.ProjectReposity {
	return &projectReposity{
		database:   db,
		collection: domain.CollectionPorject,
	}
}

func (p *projectReposity) Create(c context.Context, project domain.Project) error {
	collection := p.database.Collection(p.collection)
	_, err := collection.InsertOne(c, project)
	return err
}
