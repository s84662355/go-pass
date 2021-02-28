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
	nsqConfig "GoPass/config/nsq"
 
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
	m.ProducerName =  fmt.Sprintf("%s:%d", m.Ip, m.Port)
	return m
}

type Message struct {
	ss       *Server
	Ip       string
	Port     int
	HostName string
	ProducerName string 
}

func (l *Message) RegisterNode() {
///	redisClient.HSet("msgRegisterNode",l.ProducerName , val)
	///redisClient.Expire(l.token, 3000*time.Second)
}

func (l *Message) Consume() {
   nsqclient.InitConsumer(nsqConfig.Config.DefaultConsumer, l.consumeHandle) 
}

func  (l *Message)   consumeHandle(message *nsq.Message) error{
		fmt.Println(message.Attempts)
		fmt.Println(string(message.Body))

		///	message.Finish()
		message.Requeue(1)
		fmt.Println("111111111")
		return nil
}

func (l *Message) StarTcpServer() {
	go l. messageRelay()
	go l.ss.Start()
}

func (l *Message) send(code string){
l.ss. SendMsg(code string, dd *protocol.TcpProtocol)
}


func (l *Message) messageRelay(){
	   nsqclient.InitConsumer(nsqConfig.Config.DefaultConsumer,l.ProducerName+"messageRelay","1", func e(message *nsq.Message) error{
                   
	   }) 
}

func (l *Message) Protocol(p *protocol.TcpProtocol) {
	fmt.Println("vdfvdfvdfvdfvd")
}

func (l *Message) Error(c string) {
	fmt.Println("3222222222dfvdfvd")
}
