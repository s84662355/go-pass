package nsq

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type NsqConsumer struct {
	Name        string   `json:"name"`
	Address     []string `json:"address"`
	Topic       string   `json:"topic"`
	Channel     string   `json:"channel"`
	Concurrency int      `json:"concurrency"`
	MaxAttempts uint16   `json:"maxAttempts"`
}
