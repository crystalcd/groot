package domain

import (
	"context"
	"time"
)

const (
	CollectionPorject = "project"
)

type Project struct {
	ProjectId   string    `bson:"project_id"`
	ProjectName string    `bson:"project_name"`
	CreateTime  time.Time `bson:"timestamp"`
}

type ProjectReposity interface {
	InsertPorject(c context.Context, project Project) error
}

type ProjectService interface {
	CreateProject(c context.Context, project Project) error
}
