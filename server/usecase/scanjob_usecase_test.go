package usecase

import (
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/internal/asyncutil"
	"github.com/panjf2000/ants/v2"
)

func TestScanJob(t *testing.T) {
	p, _ := ants.NewPool(20)
	asyncutil.Setup(p)
	scanjob := NewScanJobUseCase(&bootstrap.Env{SubfinderPath: "/Users/byronchen/go/bin/subfinder", NaabuPath: "/Users/byronchen/go/bin/naabu", HttpxPath: "/Users/byronchen/go/bin/httpx"})
	scanjob.Scan("baidu.com")
}
