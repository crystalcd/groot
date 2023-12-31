package service_test

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/domain"
	"github.com/crystal/groot/repository"
	"github.com/crystal/groot/service"
)

var App bootstrap.Application

func TestMain(m *testing.M) {
	App = bootstrap.App()
	bootstrap.InjectBeans(App.Env)
	m.Run()
}

func TestScan(t *testing.T) {
}

func TestHttpScan(t *testing.T) {
	db := App.Mongo.Database("groot")
	sr := repository.NewSubdomainRepository(db, "domain")
	tr := repository.NewTaskRepository(db)
	scanService := service.NewScanService(sr, tr)
	portMap := make(map[string][]int)
	portMap["zoom.us"] = []int{443, 80}
	scanService.BatchHttpx(portMap)
}

func TestHttpRaw(t *testing.T) {
	rawReq := `GET /recaptcha/enterprise/anchor?ar=1&k=6Lf2C54aAAAAAOOpnJT1sg39rowHN362Zj2QSyls&co=aHR0cHM6Ly96b29tZGV2LnVzOjQ0Mw..&hl=en&v=pCoGBhjs9s8EhFOHJFe8cqis&theme=light&size=invisible&cb=rtqz667khhay HTTP/1.1
Host: www.google.com
Cookie: _GRECAPTCHA=09AOJJLKseFm7Ii4_cxb1ZEWgqr4MBs0GnMYTy6GR5AtFSCCWjR37Zz8wyyyfYhYFVlynGBr8V49rkfB-N0lk-jG4; NID=511=FoUrkRqqns87HgEwM9p5h391f1q0eX2MdrFDHtOoRt870-xx3gDGjXNmWy4m4povRLg-u1YxyVVP78_bgmQ0MBNnjobXhS1DY3YADlITsmrGDqwgtquCA6x7SiXOxKU_rcocoDOK7MKo6R3q-W3abFaoj0JQazIuh7yK_26g2IM
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/115.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.5
Accept-Encoding: gzip, deflate
Referer: https://zoomdev.us/
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: iframe
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: cross-site
Te: trailers

`
	fmt.Println(rawReq)
	// not support HTTP/2
	request, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawReq)))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", request)
}

func TestScanService_Scan(t *testing.T) {
	type fields struct {
		SubdomainRepository domain.SubdomainRepository
		TaskRepository      domain.TaskRepository
	}
	type args struct {
		project domain.Project
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service.ScanService{
				SubdomainRepository: tt.fields.SubdomainRepository,
				TaskRepository:      tt.fields.TaskRepository,
			}
			s.Scan(tt.args.project)
		})
	}
}
