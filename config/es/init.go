package es

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
)

var Config = config{
	Default: "default",
	Conns:   map[string]Conn{},
}

func init() {
	content, err := ioutil.ReadFile("env/es.json")
	if err != nil {
		//log.Fatal(err)
		return
	}

	ccc := config{}
	json.Unmarshal(content, &ccc)

	if ccc.Default != "" {
		Config.Default = ccc.Default
	}

	for key, v := range ccc.Conns {
		Config.Conns[key] = v
	}

}
