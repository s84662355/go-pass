package main

import (
	_ "GoPass/lib/protocol"
	"GoPass/tcp"
	_ "fmt"
	_ "github.com/golang/protobuf/proto"
	"log"
	_ "net"
)

func main() {

	m := tcp.InitMessage("127.0.0.1", 2222, 10)
	m.StarTcpServer()
	//sss := tcp.InitServer()

	//sss.Start()

	/*

		go func() {
			addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2222")

			listener, _ := net.ListenTCP("tcp", addr)
			for {
				//循环接入所有客户端得到专线连接
				conn, _ := listener.AcceptTCP()

				log.Printf(" 88888888-dC")

				ff := tcp.InitTcpEvent(conn)

				ff.ReadFunc(func(ll *protocol.TcpProtocol) {
					log.Printf(" [vsvvfdvdfvdfvdC")

					fmt.Println(ll.GetMachineCode())
					ttt := protocol.TTTT{
						Ssss: "鸡蛋均[嘿哈]为欧[嘿[嘿哈]哈]狄沃",
					}
					responseByte, _ := proto.Marshal(&ttt)
					dd := protocol.TcpProtocol{
						MachineCode: "的辅导辅导辅导s",
						BodyBytes:   responseByte,
					}
					ff.Send(&dd)
				})

				ff.ErrorFunc(func() {
					log.Printf("1111111111")
				})

				go ff.Receive()

			}
		}()

	*/

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
