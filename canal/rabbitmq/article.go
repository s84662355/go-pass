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

func RunArticle(conn *amqp.Connection, count int32) {
	//var i int32 = 0
	ccc := make(chan int, count)
	for true {
		ch, err := conn.Channel()
		if err != nil {
			panic(err)
		}
		ccc <- 1
		go article(ch, &ccc)
		//failOnError(err, "Failed to open a channel")
	}
}

func article(ch *amqp.Channel, ccc *chan int) {
	defer func() {
		<-*ccc
		ch.Close()
		if r := recover(); r != nil {
			fmt.Println(" article err: ", r)
		}
	}()

	// 监听队列

	q, err := ch.QueueDeclare(
		"article", // 队列名称
		true,      // 是否持久化
		false,     // 是否自动删除
		false,     // 是否独立
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

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		Mtype := gjson.Get(string(d.Body), "type").String()
		switch Mtype {
		case "INSERT":
			doArticle(string(d.Body))
			break
		case "UPDATE":
			doArticle(string(d.Body))
			break
		case "DELETE":
			break
		}

	}

}

func doArticle(body string) {
	for _, res := range gjson.Get(body, "data").Array() {
		res.ForEach(func(key, value gjson.Result) bool {
			if key.String() == "id" {
				model.Article{}.PostDataById(value.String())
			}
			return true
		})
	}
}
