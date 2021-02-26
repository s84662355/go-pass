package tcp

import (
	"GoPass/lib/protocol"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
	"sync"
)

func InitTcpEvent(conn *net.TCPConn) *TcpEvent {
	e := &TcpEvent{}
	e.conn = conn
	e.m = new(sync.RWMutex)
	e.readBuf = newBuffer(conn, BufLen)
	e.writeBuf = make([]byte, BufLen)

	return e
}

type TcpEvent struct {
	conn      *net.TCPConn
	m         *sync.RWMutex
	readBuf   buffer // 读缓冲
	writeBuf  []byte
	readFunc  func(*protocol.TcpProtocol)
	Code      string
	errorFunc func(*TcpEvent)
}

func (l *TcpEvent) ReadFunc(readFunc func(*protocol.TcpProtocol)) {
	l.readFunc = readFunc
}

func (l *TcpEvent) ErrorFunc(errorFunc func(*TcpEvent)) {
	l.errorFunc = errorFunc
}

func (l *TcpEvent) Send(response *protocol.TcpProtocol) error {
	defer l.m.Unlock()
	l.m.Lock()
	responseByte, _ := proto.Marshal(response)
	fmt.Println(responseByte)
	responseLen := uint16(len(responseByte))
	fmt.Println(responseLen)

	binary.BigEndian.PutUint16(l.writeBuf[0:TcpProtocolLen], responseLen)
	copy(l.writeBuf[TcpProtocolLen:], responseByte[:responseLen])
	_, e := l.conn.Write(l.writeBuf[0 : TcpProtocolLen+responseLen])

	return e

}

// Decode 解码数据
func (c *TcpEvent) decode() (*protocol.TcpProtocol, bool) {
	var err error
	// 读取数据类型
	typeBuf, err := c.readBuf.seek(0, TcpProtocolLen)
	if err != nil {
		return nil, false
	}

	tcpRequestHeadLen := int(binary.BigEndian.Uint16(typeBuf))

	// 读取数据长度
	lenBuf, err := c.readBuf.read(TcpProtocolLen, tcpRequestHeadLen)
	if err != nil {
		return nil, false
	}

	head := &protocol.TcpProtocol{}

	proto.Unmarshal(lenBuf, head)

	return head, true
}

func (c *TcpEvent) Release() {
	c.conn.Close()
}

func (c *TcpEvent) Receive() {

	for {
		_, err := c.readBuf.readFromReader()
		if err != nil {
			fmt.Println(err)
			c.errorFunc(c)
			return
		}

		for {
			head, ok := c.decode()
			if ok {
				head.ECode = c.Code
				c.readFunc(head)
				//	_ = c.HandlePackage(pack)
				continue
			}
			break
		}
	}
}
