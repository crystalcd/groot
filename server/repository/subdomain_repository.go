package repository

import (
	"context"

	"github.com/crystal/groot/domain"
	"github.com/qiniu/qmgo"
)

type subdomainRepository struct {
	database   *qmgo.Database
	collection string
}

func NewSubdomainRepository(db *qmgo.Database, collection string) domain.SubdomainRepository {
	return &subdomainRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *subdomainRepository) InsertSubdomains(c context.Context, subdomains []domain.Subdomain) error {
	collection := sr.database.Collection(sr.collection)
	_, err := collection.InsertMany(c, subdomains)
	return err
}
