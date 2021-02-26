package nsq

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type NsqProducer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
