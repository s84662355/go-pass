package main

import (
	"GoPass/lib/protocol"
	"GoPass/tcp"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2222")

	conn, _ := net.DialTCP("tcp", nil, addr)

	ff := tcp.InitTcpEvent(conn)

	dd := protocol.TcpProtocol{
		//MachineCode: "gngh回话mhsdcs",
		BodyBytes: []byte{3, 64, 3, 5, 45},
	}
	ff.Send(&dd)

	ff.ReadFunc(func(ll *protocol.TcpProtocol) {
		log.Printf(" 322223让3热3vdfvdC")
		head := protocol.TTTT{}
		proto.Unmarshal(ll.GetBodyBytes(), &head)

		fmt.Println(head.GetSsss())

	})

	go ff.Receive()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
