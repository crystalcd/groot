package domain

import "time"

type Test struct {
	Id          string    `bson:"id" json:"id"`
	Name        string    `json:"name"`
	ProjectId   string    `bson:"project_id" json:"project_id"`
	ProjectName string    `bson:"project_name" json:"project_name"`
	Domains     []string  `bson:"domains" json:"domains"`
	CreateTime  time.Time `bson:"timestamp" json:"timestamp"`
}
