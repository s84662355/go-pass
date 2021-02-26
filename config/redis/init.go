package redis

import (
	"encoding/json"
	_ "fmt"
	_ "github.com/go-redis/redis"
	_ "github.com/tidwall/gjson"
	"io/ioutil"
)

var Config = config{
	Default: "default",
	OptionsConns: map[string]Options{
		"default": {
			Addr:     "127.0.0.1:6379",
			Password: "",
			Db:       0,
			PoolSize: 500,
		},
	},
}

func init() {
	content, err := ioutil.ReadFile("env/redis.json")
	if err != nil {
		//log.Fatal(err)
		return
	}

	ccc := config{}
	json.Unmarshal(content, &ccc)

	if ccc.Default != "" {
		Config.Default = ccc.Default
	}

	for key, v := range ccc.OptionsConns {
		Config.OptionsConns[key] = v
	}

}
