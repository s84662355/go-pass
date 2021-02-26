package redis

import (
	_ "encoding/json"
	_ "fmt"
	_ "github.com/go-redis/redis"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type config struct {
	OptionsConns map[string]Options `json:"options_conns"`
	Default      string             `json:"default"`
}
