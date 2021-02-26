package rabbitmq

import (
	"GoPass/es/model"
	_ "GoPass/lib/es"
	_ "GoPass/lib/mysql"
	_ "GoPass/lib/rabbitmq"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
	"log"
)

func RunCate(conn *amqp.Connection, count int32) {
	//var i int32 = 0
	ccc := make(chan int, count)
	for true {
		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}
		ccc <- 1
		go cate(ch, &ccc)
		//failOnError(err, "Failed to open a channel")
	}

}

func cate(ch *amqp.Channel, ccc *chan int) {

	defer func() {
		<-*ccc
		ch.Close()
		if r := recover(); r != nil {
			fmt.Println(" article err: ", r)
		}
	}()
	// 监听队列
	q, err := ch.QueueDeclare(
		"cate", // 队列名称
		true,   // 是否持久化
		false,  // 是否自动删除
		false,  // 是否独立
		false, nil,
	)

	if err != nil {
		panic(err)
	}

	var tt map[string]interface{} = nil
	//ch.QueueUnbind(q.Name, q.Name, "canal", tt)

	err = ch.QueueBind(q.Name, q.Name, "canal", false, tt)

	if err != nil {
		panic(err)
	}
	// 消费队列
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}
	// 申明一个goroutine,一遍程序始终监听

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)

		Mtype := gjson.Get(string(d.Body), "type").String()

		switch Mtype {
		case "INSERT":
			doCate(string(d.Body))
			break
		case "UPDATE":
			doCate(string(d.Body))
			break
		case "DELETE":
			break
		}

	}

}

func doCate(body string) {
	for _, res := range gjson.Get(body, "data").Array() {
		res.ForEach(func(key, value gjson.Result) bool {
			if key.String() == "id" {
				ammmr := model.Article{}.GetByCateId(value.String())
				for _, rrr := range ammmr {
					rrr.Update()
				}
			}
			return true
		})
	}
}
