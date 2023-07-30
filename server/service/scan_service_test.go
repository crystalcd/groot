package service_test

import (
	"testing"

	"github.com/crystal/groot/bootstrap"
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
	scanService := service.NewScanService(s, n, h)
	scanService.Scan("zoom", "zoom.us")
}
