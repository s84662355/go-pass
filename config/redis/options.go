package redis

import (
	_ "encoding/json"
	_ "fmt"
	_ "github.com/go-redis/redis"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type Options struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}
