package domain

import "sync"

type Result struct {
	sync.RWMutex
	R map[string][]string
}

func (r *Result) Set(key string, value []string) {
	r.Lock()
	defer r.Unlock()
	r.R[key] = value
}

func (r *Result) Get(key string) []string {
	r.RLock()
	defer r.RUnlock()
	return r.R[key]
}
