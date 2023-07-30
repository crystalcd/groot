package service_test

import (
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/repository"
	"github.com/crystal/groot/service"
	"github.com/crystal/groot/tools/scan"
)

func TestMain(m *testing.M) {
	app := bootstrap.App()
	bootstrap.InjectBeans(app.Env)
	m.Run()
}

func TestScan(t *testing.T) {
	s := scan.NewSubfinder()
	n := scan.NewNaabu()
	h := scan.NewHttpx()
	sr := repository.NewSubdomainRepository(bootstrap.App().Mongo.Database("groot"), "domain")
	scanService := service.NewScanService(s, n, h, sr)
	scanService.Scan("zoom", "zoom.us")
}

func TestHttpScan(t *testing.T) {
	s := scan.NewSubfinder()
	n := scan.NewNaabu()
	h := scan.NewHttpx()
	sr := repository.NewSubdomainRepository(bootstrap.App().Mongo.Database("groot"), "domain")
	scanService := service.NewScanService(s, n, h, sr)
	portMap := make(map[string][]int)
	portMap["zoom.us"] = []int{443, 80}
	scanService.BatchHttpx(portMap)
}
