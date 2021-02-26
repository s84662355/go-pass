package rabbitmq

import (
	rabbitmqConfig "GoPass/config/rabbitmq"
	_ "context"
	_ "encoding/json"
	_ "fmt"
	"github.com/streadway/amqp"
	"sync"
)

var amqpDatabases sync.Map

func InitAmqpConnect() {
	for k, v := range rabbitmqConfig.Config.Conns {
		ConnectAmqp(v, k)
	}
}

func ConnectAmqp(cc rabbitmqConfig.Conn, name string) *amqp.Connection {

	conn, err := amqp.Dial(cc.Url)
	if err != nil {
		panic(err)
	}
	amqpDatabases.Store(name, conn)
	return conn
}

func Amqp(name string) *amqp.Connection {
	client, _ := amqpDatabases.Load(name)
	return client.(*amqp.Connection)
}

func DisconnectAmqp() {
	amqpDatabases.Range(func(key, value interface{}) bool {
		defer value.(*amqp.Connection).Close()
		return true
	})
}
