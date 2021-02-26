package canal

import (
	canalConfig "GoPass/config/canal"
	"fmt"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/withlin/canal-go/client"
	"sync"
	_ "time"
)

var canalConnsMap sync.Map

func InitCanal() {
	for k, v := range canalConfig.Config.Conns {

		ConnectCanal(v, k)
		fmt.Println(v)
		fmt.Println(k)
		//	return
	}
}

func ConnectCanal(config canalConfig.Conn, name string) *client.SimpleCanalConnector {
	connector := client.NewSimpleCanalConnector(config.Address, config.Port, config.Username, config.Password, config.Destination, config.SoTimeOut, config.IdleTimeOut)
	connector.Connect()
	connector.Subscribe(config.Subscribe)
	canalConnsMap.Store(name, connector)
	return connector
}

func Canal(name string) *client.SimpleCanalConnector {
	connector, _ := canalConnsMap.Load(name)
	return connector.(*client.SimpleCanalConnector)
}

func DisconnectCanal() {
	canalConnsMap.Range(func(key, value interface{}) bool {
		defer value.(*client.SimpleCanalConnector).DisConnection()
		return true
	})
}
