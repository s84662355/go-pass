package session

import (
	"GoPass/lib/redis"
	_ "encoding/json"
	"time"
)

var redisClient = redis.GetRedis()

type HttpSessionInterface interface {
	init(token string) error
	Heartbeat(second int64)
	Set(key string, val interface{}) bool
	Get(key string) interface{}
	Del()
}

type httpSession struct {
	token string
	data  map[string]interface{}
}

func CreateHttpSession(token string) httpSession {
	c := httpSession{}
	c.init(token)
	return c
}

func (l *httpSession) init(token string) error {
	l.token = token
	l.data = make(map[string]interface{})
	//redisClient.Exists(l.token)
	val, err := redisClient.HGetAll(l.token).Result()
	if err != nil {
		return err
	}
	for i, v := range val {
		l.data[i] = v
	}
	go redisClient.Expire(l.token, 3000*time.Second)
	return nil
}

func (l *httpSession) Heartbeat(second int64) {
	go redisClient.Expire(l.token, time.Duration(second)*time.Second)
}

func (l *httpSession) Set(key string, val interface{}) bool {
	l.data[key] = val
	if !redisClient.HSet(l.token, key, val).Val() {
		delete(l.data, key)
		return false
	}
	return true
}

func (l *httpSession) Get(key string) interface{} {
	if val, ok := l.data[key]; ok {
		return val
	} else {
		return nil
	}
}

func (l *httpSession) Del() {
	go redisClient.Del(l.token)
}
