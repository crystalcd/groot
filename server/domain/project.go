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
	Domains     []string  `bson:"domains"`
	CreateTime  time.Time `bson:"timestamp"`
}

type ProjectReposity interface {
	Create(c context.Context, project Project) error
}

type ProjectService interface {
	CreateProject(c context.Context, project Project) error
}
