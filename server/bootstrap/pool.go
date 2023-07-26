package bootstrap

import (
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

func NewPool(env *Env, *logrus.Logger) *ants.Pool {
	pool, err := ants.NewPool(env.AsyncPoolCount)
	return pool
}
