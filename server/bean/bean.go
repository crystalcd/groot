package bean

import "sync"

type Config struct {
	Path string
}

type Param struct {
	Target  string
	Project string
}

type Result struct {
	sync.RWMutex
	DomainResult map[string][]string
}

func (r *Result) SetSubDomain(domain, subdoamin string) {
	r.Lock()
	defer r.Unlock()
	sudomains, ok := r.DomainResult[domain]
	if ok {
		sudomains = append(sudomains, subdoamin)
	} else {
		sudomains = []string{domain}
	}
	r.DomainResult[domain] = sudomains
}

type Domain struct {
	Project string `bson:"project,omitempty"`
	Domain  string `bson:"domain,omitempty"`
	From    string `bson:"from,omitempty"`
}

type AddProject struct {
	Domains     string `json:"domains"`
	ProjectName string `json:"name"`
}
