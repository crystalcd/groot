package eventbus

import (
	"sync"

	"github.com/crystal/groot/logging"
)

var EB = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

type DataEvent struct {
	Data  interface{}
	Topic string
}

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	logging.RuntimeLog.Infof("<--------------publish topic: %s", topic)
	eb.rm.RLock()
	defer eb.rm.RUnlock()
	if chans, found := eb.subscribers[topic]; found {
		// 这样做是因为切片引用相同的数组，即使它们是按值传递的
		// 因此我们正在使用我们的元素创建一个新切片，从而能正确地保持锁定
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlice DataChannelSlice) {
			for _, ch := range dataChannelSlice {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	logging.RuntimeLog.Infof("-------------->subscribe topic: %s", topic)
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
}
