package mysql

import (
	"fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

type Conn struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       string `json:"db"`
}

func (l Conn) GetDsn() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", l.Username, l.Password, l.Host, l.Port, l.Db)
}
