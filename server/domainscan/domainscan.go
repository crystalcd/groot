package domainscan

import (
	"context"
	"reflect"
	"strings"
	"sync"

	"github.com/crystal/groot/bean"
	"github.com/crystal/groot/db"
	"github.com/crystal/groot/eventbus"
	"github.com/crystal/groot/logging"
	"github.com/crystal/groot/pool"
)

func init() {
	listenSubfinder := make(chan eventbus.DataEvent)
	listenAssetfinder := make(chan eventbus.DataEvent)
	eventbus.EB.Subscribe(TopicSubfinder, listenSubfinder)
	eventbus.EB.Subscribe(TopicAssetfinder, listenAssetfinder)
	go func() {
		for {
			select {
			case d := <-listenSubfinder:
				go printDataEvent(TopicSubfinder, d)
			case d := <-listenAssetfinder:
				go printDataEvent(TopicAssetfinder, d)
			}
		}
	}()
}

func printDataEvent(topic string, data eventbus.DataEvent) {
	switch v := data.Data.(type) {
	case Subfinder:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.Param.Target)
	case Assetfinder:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.Param.Target)
	case *Subfinder:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.Param.Target)
	default:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, "default")
	}

}

const TopicSubfinder = "topic_subfinder"
const TopicAssetfinder = "topic_assetfinder"

type Scan interface {
	Scan(scan Scan)
	AsyncScan(scan Scan)
	run(domain string)
	ParseResult(domain string, data []byte)
	Write2MongoDB(scan Scan)
}

type DomainScan struct {
	Config bean.Config
	Param  bean.Param
	Result bean.Result
}

func (d *DomainScan) Scan(scan Scan) {
	logging.RuntimeLog.Info("---------------scan start-----------")
	d.Result.DomainResult = map[string][]string{}
	var wg sync.WaitGroup
	for _, line := range strings.Split(d.Param.Target, ",") {
		wg.Add(1)
		domain := strings.TrimSpace(line)
		pool.DOMAIN_SCAN.Submit(func() {
			defer wg.Done()
			scan.run(domain)
		})
	}
	wg.Wait()
	eventbus.EB.Publish(TopicSubfinder, scan)
}

func (d *DomainScan) AsyncScan(scan Scan) {
	pool.DOMAIN_SCAN.Submit(func() {
		scan.Scan(scan)
	})
}

func (d *DomainScan) run(domain string) {
	logging.RuntimeLog.Warn("you need overriding the method Run()")
}

func (d *DomainScan) ParseResult(domain string, data []byte) {
	for _, line := range strings.Split(string(data), "\n") {
		subdomain := strings.TrimSpace(line)
		if subdomain == "" {
			continue
		}
		d.Result.SetSubDomain(domain, subdomain)
	}
}

func (d *DomainScan) Write2MongoDB(scan Scan) {
	t := reflect.TypeOf(scan)
	from := t.Name()
	domainMap := d.Result.DomainResult
	alldomains := []string{}
	for key, value := range domainMap {
		alldomains = append(alldomains, key)
		alldomains = append(alldomains, value...)
	}
	allObjs := []bean.Domain{}
	for _, line := range alldomains {
		domainObj := bean.Domain{
			Project: d.Param.Project,
			Domain:  line,
			From:    from,
		}
		allObjs = append(allObjs, domainObj)
	}
	db.DomainCli.InsertMany(context.Background(), allObjs)
}
