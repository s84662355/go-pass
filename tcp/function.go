package tcp

import (
	_ "GoPass/lib/helper"
	_ "GoPass/lib/protocol"
	"GoPass/lib/redis"
	_ "encoding/binary"
	_ "fmt"
	_ "github.com/golang/protobuf/proto"
	_ "net"
	_ "sync"
	_ "time"
)

var redisClient = redis.GetRedis()

func registerNode(hostname string) {

}
