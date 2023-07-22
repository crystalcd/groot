package domain

import "context"

const (
	CollectionSubdomains = "domains"
)

type Subdomain struct {
	Project string `bson:"project,omitempty"`
	Domain  string `bson:"domain,omitempty"`
	From    string `bson:"from,omitempty"`
}

type SubdomainRepository interface {
	InsertSubdomains(c context.Context, subdomains []Subdomain) error
}
