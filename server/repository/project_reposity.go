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

func NewProjectReposity(db *qmgo.Database, collection string) domain.ProjectReposity {
	return &projectReposity{
		database:   db,
		collection: collection,
	}
}

func (p *projectReposity) InsertPorject(c context.Context, project domain.Project) error {
	collection := p.database.Collection(p.collection)
	_, err := collection.InsertOne(c, project)
	return err
}
