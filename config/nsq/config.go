package nsq

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

var Config = config{
	DefaultConsumer: "default",
	DefaultProducer: "default",
	NsqConsumer: map[string]NsqConsumer{
		"default": {
			Name:        "default",
			Address:     []string{"127.0.0.1:4161"},
			Topic:       "test",
			Channel:     "1",
			Concurrency: 10,
			MaxAttempts: 3,
		},
	},
	NsqProducer: map[string]NsqProducer{
		"default": {
			Name:    "default",
			Address: "127.0.0.1:4150",
		},
	},
}

type config struct {
	NsqConsumer     map[string]NsqConsumer `json:"nsq_consumer"`
	NsqProducer     map[string]NsqProducer `json:"nsq_producer"`
	DefaultConsumer string                 `json:"default_consumer"`
	DefaultProducer string                 `json:"default_producer"`
}
