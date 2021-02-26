package tcp

import (
	"GoPass/lib/helper"
	"GoPass/lib/protocol"
	_ "encoding/binary"
	"fmt"
	_ "github.com/golang/protobuf/proto"
	"net"
	"sync"
	_ "time"
)

func InitServer() *Server {
	ss := Server{}
	ss.m = &sync.RWMutex{}
	ss.MaxConnLen = 1000
	ss.Ip = "127.0.0.1"
	ss.Port = 2222
	ss.AcceptCount = 10

	return &ss

}

type Handler interface {
	// 方法列表
	Protocol(*protocol.TcpProtocol)
	Error(string)
}

type Server struct {
	MaxConnLen  uint64
	m           *sync.RWMutex
	tcpEventMap sync.Map
	connLen     uint64
	Ip          string
	Port        int
	listener    *net.TCPListener
	H           Handler
	AcceptCount uint16 // 接收建立连接的groutine数量
}

func (l *Server) Start() {

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", l.Ip, l.Port))
	l.listener, _ = net.ListenTCP("tcp", addr)
	var i uint16 = 0
	for ; i < l.AcceptCount; i++ {
		go l.accept()
	}
}

func (l *Server) accept() {
	for {
		conn, _ := l.listener.AcceptTCP()
		err := conn.SetKeepAlive(true)
		if err != nil {
		}

		ff := InitTcpEvent(conn)
		ff.Code = helper.GetRandomBoth(10)
		ff.ReadFunc(l.readFunc)
		ff.ErrorFunc(l.errorFunc)
		l.tcpEventMap.Store(ff.Code, ff)
		go ff.Receive()

	}
}

func (l *Server) SendMsg(code string, dd *protocol.TcpProtocol) {
	v, ok := l.tcpEventMap.Load(code)
	if ok {
		v.(*TcpEvent).Send(dd)
	}
}

func (l *Server) Close(code string) {
	v, ok := l.tcpEventMap.Load(code)
	if ok {
		v.(*TcpEvent).Release()
		l.tcpEventMap.Delete(code)
	}
}

func (l *Server) readFunc(ll *protocol.TcpProtocol) {
	l.H.Protocol(ll)
}

func (l *Server) errorFunc(ll *TcpEvent) {
	l.H.Error(ll.Code)
	ll.Release()
	l.tcpEventMap.Delete(ll.Code)

}
