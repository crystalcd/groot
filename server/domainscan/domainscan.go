package domainscan

import (
	"github.com/crystal/groot/eventbus"
	"github.com/crystal/groot/logging"
)

const TopicSubfinder = "topic_subfinder"
const TopicAssetfinder = "topic_assetfinder"

type DomainScan interface {
	Do()
	AsyncDo()
	Run(domain string)
	ParseResult(domain string, data []byte)
	Write2MongoDB()
}

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
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.S.Param.Target)
	case Assetfinder:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.S.Param.Target)
	case *Subfinder:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, v.S.Param.Target)
	default:
		logging.RuntimeLog.Warnf("%s ended type %s", topic, "default")
	}

}
