package eventbus

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func publisTo(topic string, data string) {
	for {
		EB.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}

func TestEventMain(t *testing.T) {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)
	EB.Subscribe("topic1", ch1)
	EB.Subscribe("topic2", ch2)
	EB.Subscribe("topic2", ch3)
	go publisTo("topic1", "Hi topic 1")
	go publisTo("topic2", "Welcome to topic 2")
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}

func TestEventBus_Subscribe(t *testing.T) {
	type fields struct {
		subscribers map[string]DataChannelSlice
		rm          sync.RWMutex
	}
	type args struct {
		topic string
		ch    DataChannel
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
			eb := &EventBus{
				subscribers: tt.fields.subscribers,
				rm:          tt.fields.rm,
			}
			eb.Subscribe(tt.args.topic, tt.args.ch)
		})
	}
}
