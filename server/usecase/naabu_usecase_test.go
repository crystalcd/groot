package usecase

import (
	"fmt"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/panjf2000/ants/v2"
)

func TestScanNaabu(t *testing.T) {
	p, _ := ants.NewPool(10000)
	asyncutil.Setup(p)
	naabu := NewNaabuUseCase(&bootstrap.Env{NaabuPath: "/Users/byronchen/go/bin/naabu"})
	rs := naabu.Scan(`baiduspider-123-125-66-50.crawl.baidu.com`)
	for k, v := range rs.R {
		fmt.Printf("key:%s value:%v", k, v)
	}
}
