package canal

import (
	_ "fmt"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

//address string, port int, username string, password string, destination string, soTimeOut int32, idleTimeOut int32

type Conn struct {
	Address     string `json:"address"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Destination string `json:"destination"`
	SoTimeOut   int32  `json:"soTimeOut"`
	IdleTimeOut int32  `json:"idleTimeOut"`
	Subscribe   string `json:"subscribe"`
}
