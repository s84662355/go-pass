package service

import (
	"GoPass/lib/helper"
	"GoPass/lib/redis"

	_ "fmt"
	_ "strconv"
)

var ReadCountQueue helper.ChanQueue
var redisClient = redis.GetRedis()

func GlobalInit() {
	ReadCountQueue = helper.ChanQueue{}.Create(10)
	ReadCountQueue.Start(func(data interface{}) {
		Article{}.UpdateReadAmount(data)
	}, 3)
}
