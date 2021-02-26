package nsqclient

import (
	nsqConfig "GoPass/config/nsq"
	_ "fmt"
	"github.com/nsqio/go-nsq"
	"sync"
)

var nsqProducerMap sync.Map

func InitProducer(name string) *nsq.Producer {

	config := nsqConfig.Config.NsqProducer[name]
	return ConnectNsqProducer(config.Address, name)

}

func GetProducer() *nsq.Producer {
	return Producer(nsqConfig.Config.DefaultProducer)
}

func ConnectNsqProducer(address string, name string) *nsq.Producer {
	producer, err := nsq.NewProducer(address, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	nsqProducerMap.Store(name, producer)
	return producer
}

func Producer(name string) *nsq.Producer {
	producer, _ := nsqProducerMap.Load(name)
	return producer.(*nsq.Producer)
}

func DisconnectProducer() {
	nsqProducerMap.Range(func(key, value interface{}) bool {
		defer value.(*nsq.Producer).Stop()
		return true
	})
}
