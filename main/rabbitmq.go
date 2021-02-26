package main

import (
	lmq "GoPass/canal/rabbitmq"
	_ "GoPass/es/model"
	"GoPass/lib/es"
	"GoPass/lib/mysql"
	"GoPass/lib/rabbitmq"
	_ "fmt"
	_ "github.com/streadway/amqp"
	_ "github.com/tidwall/gjson"
	"log"
)

func main() {
	rabbitmq.InitAmqpConnect()
	mysql.InitMysqlConnect()
	es.InitEsConnect()

	go func() {
		conn := rabbitmq.Amqp("default")
		lmq.RunArticle(conn, 10)
	}()

	go func() {
		conn := rabbitmq.Amqp("default")
		lmq.RunCate(conn, 10)
	}()
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

// 帮助函数检测每一个amqp调用
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
