package rabbitmq

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type Conn struct {
	Url string `json:"url"`
}
