package canal

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
)

//"127.0.0.1", 11111, "", "", "example", 60000, 60*60*1000
var Config = config{
	Default: "default",
	Conns: map[string]Conn{
		"default": {
			Address:     "127.0.0.1",
			Port:        11111,
			Username:    "",
			Password:    "",
			Destination: "example",
			SoTimeOut:   60000,
			IdleTimeOut: 60 * 60 * 1000,
		},
	},
}

func init() {
	content, err := ioutil.ReadFile("env/canal.json")
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
