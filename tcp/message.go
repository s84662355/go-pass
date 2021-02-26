package tcp

import (
	_ "GoPass/lib/helper"
	"GoPass/lib/nsqclient"
	"GoPass/lib/protocol"
	"GoPass/lib/redis"
	_ "encoding/binary"
	"fmt"
	_ "github.com/golang/protobuf/proto"
	_ "net"
	_ "sync"
	_ "time"
)

var redisClient = redis.GetRedis()

func InitMessage(ip string, port int, acceptCount uint16) *Message {
	m := &Message{}
	m.ss = InitServer()
	m.ss.Ip = ip
	m.ss.Port = port
	m.ss.AcceptCount = acceptCount
	m.ss.H = m
	m.Ip = ip
	m.Port = port
	return m
}

type Message struct {
	ss       *Server
	Ip       string
	Port     int
	HostName string
}

func (l *Message) RegisterNode() {
	///redisClient.Expire(l.token, 3000*time.Second)
}

func (l *Message) Consume() {

}

func (l *Message) StarTcpServer() {
	go l.ss.Start()
}

func (l *Message) Protocol(p *protocol.TcpProtocol) {
	fmt.Println("vdfvdfvdfvdfvd")
}

func (l *Message) Error(c string) {
	fmt.Println("3222222222dfvdfvd")
}
