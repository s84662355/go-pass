package mysql

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
)

var Config = config{
	Default: "default",
	Conns: map[string]Conn{
		"default": {
			Username: "root",
			Password: "123456",
			Host:     "127.0.0.1",
			Port:     3306,
			Db:       "test",
		},
	},
}

func init() {
	content, err := ioutil.ReadFile("env/mysql.json")
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
