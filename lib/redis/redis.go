package redis

import (
	_ "GoPass/config"
	redisConfig "GoPass/config/redis"
	"github.com/go-redis/redis"
	"sync"
)

var redisDatabases sync.Map

func init() {

	for k, v := range redisConfig.Config.OptionsConns {
		client := redis.NewClient(&redis.Options{
			Addr:     v.Addr,
			Password: v.Password,
			DB:       v.Db,
			PoolSize: v.PoolSize,
		})
		redisDatabases.Store(k, client)
	}

	/*
		client := redis.NewClient(&redis.Options{
			Addr:     config.RedisAddr,
			Password: config.RedisPassword,
			DB:       config.RedisDb,
			PoolSize: config.RedisPoolSize,
		})
		redisDatabases.Store("default", client)


	*/

}

func GetRedis() *redis.Client {
	value, ok := redisDatabases.Load(redisConfig.Config.Default)
	if ok {
		return value.(*redis.Client)
	}
	panic("failed to connect redis database:" + redisConfig.Config.Default)
}

func GetRedisClient(name string) *redis.Client {
	value, ok := redisDatabases.Load(name)
	if ok {
		return value.(*redis.Client)
	}
	panic("failed to connect redis database:" + name)
}
