package usecase

import (
	"fmt"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/panjf2000/ants/v2"
)

func TestScan(t *testing.T) {
	p, _ := ants.NewPool(100)
	asyncutil.Setup(p)
	subfinder := NewSubfinderUseCase(&bootstrap.Env{SubfinderPath: "/Users/byronchen/go/bin/subfinder"})
	rs := subfinder.Scan("baidu.com,zoom.us")
	for k, v := range rs.R {
		fmt.Printf("key:%s value:%v", k, v)
	}
}
