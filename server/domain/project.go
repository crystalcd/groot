package domain

import (
	"context"
	"time"
)

const (
	CollectionPorject = "project"
)

type Project struct {
	ProjectName string    `bson:"project_name" json:"projectName"`
	Domains     []string  `bson:"domains" json:"domains"`
	CreateTime  time.Time `bson:"timestamp"`
	OwnId       string    `bson:"own_id"`
	Version     string    `bson:"version"`
}

type ProjectRepository interface {
	Create(c context.Context, project Project) error
	QueryByName(c context.Context, name string) ([]Project, error)
}

type ProjectService interface {
	CreateProject(c context.Context, project Project) error
}
