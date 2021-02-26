package canal

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type config struct {
	Conns   map[string]Conn `json:"conns"`
	Default string          `json:"default"`
}
