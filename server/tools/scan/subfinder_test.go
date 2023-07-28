package scan_test

import (
	"log"
	"sync"
	"testing"

	"github.com/crystal/groot/bootstrap"
	"github.com/crystal/groot/tools/scan"
)

func TestMain(m *testing.M) {
	app := bootstrap.App()
	bootstrap.InjectBeans(app.Env)
	m.Run()
}

func TestScan(t *testing.T) {
	subfinder := scan.NewSubfinder()
	subdomains, err := subfinder.Scan("baidu.com")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("subdomains: %v", subdomains)
}

func TestNaabu(t *testing.T) {
	naabu := scan.NewNaabu()
	ports, err := naabu.Scan("baidu.com")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("ports: %v", ports)
}

func TestHttpx(t *testing.T) {
	httpx := scan.NewHttpx()
	ports := []string{"80", "443"}
	rs, err := httpx.Scan("baidu.com", ports)
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	bootstrap.Logger.Infof(" length %d, %+v", len(rs), rs)
}

func TestDomainAndPort(t *testing.T) {
	domainPort := make(map[string][]string)
	var l sync.RWMutex
	subfinder := scan.NewSubfinder()
	naabu := scan.NewNaabu()
	subdomains, err := subfinder.Scan("baidu.com")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	var wg sync.WaitGroup
	for _, sudomain := range subdomains {
		host := sudomain
		wg.Add(1)
		go func() {
			ports, err := naabu.Scan(host)
			if err != nil {
				bootstrap.Logger.Error(err)
				//ignore
			}
			l.Lock()
			domainPort[host] = ports
			l.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	bootstrap.Logger.Infof("%+v", domainPort)
}

func TestWaybackUrls(t *testing.T) {
	waybackurls := scan.NewWaybackurls()
	rs, err := waybackurls.Scan("zoom.us")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	bootstrap.Logger.Info(rs)
}

func TestWappalyze(t *testing.T) {
	wappalyzer := scan.NewWappalyze()
	rs, err := wappalyzer.Scan("http://zoom.us")
	if err != nil {
		bootstrap.Logger.Error(err)
	}
	bootstrap.Logger.Infof("%+v", rs)
}
