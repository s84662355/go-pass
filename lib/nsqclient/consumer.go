package nsqclient

import (
	nsqConfig "GoPass/config/nsq"
	_ "fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"time"
)

var nsqConsumerMap sync.Map

func InitConsumer(name string, handle func(message *nsq.Message) error) *nsq.Consumer {

	config := nsqConfig.Config.NsqConsumer[name]
	return NsqConsumer(name, config.Topic, config.Channel, config.Address, handle, config.Concurrency, config.MaxAttempts)

}

func GetConsumer() *nsq.Consumer {
	return Consumer(nsqConfig.Config.DefaultConsumer)
}

// nsqConsumer 消费消息
func NsqConsumer(name, topic, channel string, hosts []string, handle func(message *nsq.Message) error, concurrency int, maxAttempts uint16) *nsq.Consumer {

	//MaxAttempts 重试次数
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = 1 * time.Second
	conf.MaxInFlight = 10 + len(hosts) //最大允许向两台NSQD服务器接受消息，默认是1，要特别注意
	conf.MaxAttempts = maxAttempts

	consumer, err := nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)
	//consumer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	err = consumer.ConnectToNSQLookupds(hosts)

	if err != nil {
		panic(err)
	}

	nsqConsumerMap.Store(name, consumer)

	return consumer
}

func Consumer(name string) *nsq.Consumer {
	consumer, _ := nsqConsumerMap.Load(name)
	return consumer.(*nsq.Consumer)
}

func DisconnectConsumer() {
	nsqConsumerMap.Range(func(key, value interface{}) bool {
		defer value.(*nsq.Consumer).Stop()
		return true
	})
}
