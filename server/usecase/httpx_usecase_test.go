package usecase

import (
	"fmt"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/panjf2000/ants/v2"
)

func TestHttpx(t *testing.T) {
	p, _ := ants.NewPool(100)
	asyncutil.Setup(p)
	httpx := NewHttpxUseCase(&bootstrap.Env{HttpxPath: "/Users/crystal/go/bin/httpx"})
	ports := []int{80, 443}
	rs := httpx.Scan("zoom.us", ports)
	for _, obj := range rs.R {
		fmt.Printf("obj=%+v \n", obj)
	}
}
