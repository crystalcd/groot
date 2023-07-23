package asyncutil

import (
	"log"

	"github.com/panjf2000/ants/v2"
)

var pool *ants.Pool

func AsyncDo(task func()) error {
	if pool == nil {
		log.Fatal("pool is nil")
	}
	return pool.Submit(task)
}

func Setup(p *ants.Pool) {
	pool = p
}
