package domain

import (
	"context"
	"time"
)

const (
	CollectionSubdomains = "domains"
)

type Subdomain struct {
	Project    string    `bson:"project,omitempty"`
	Domain     string    `bson:"domain,omitempty"`
	From       string    `bson:"from,omitempty"`
	Ports      []int     `bson:"ports"`
	CreateTime time.Time `bson:"create_time"`
}

type SubdomainRepository interface {
	InsertSubdomains(c context.Context, subdomains []Subdomain) error
}
