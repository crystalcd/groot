package usecase

import (
	"fmt"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/panjf2000/ants/v2"
)

func TestScanNaabu(t *testing.T) {
	p, _ := ants.NewPool(100)
	asyncutil.Setup(p)
	naabu := NewNaabuUseCase(&bootstrap.Env{NaabuPath: "/Users/byronchen/go/bin/naabu"})
	rs := naabu.Scan("zoom.us")
	for k, v := range rs.R {
		fmt.Printf("key:%s value:%v", k, v)
	}
}
