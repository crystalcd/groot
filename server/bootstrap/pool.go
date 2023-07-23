package bootstrap

import (
	"github.com/panjf2000/ants/v2"
)

func NewPool(env *Env) *ants.Pool {
	pool, _ := ants.NewPool(env.AsyncPoolCount)
	return pool
}
