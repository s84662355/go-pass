package es

import (
	_ "fmt"
	esc "github.com/olivere/elastic/v7/config"
	_ "github.com/tidwall/gjson"
	_ "io/ioutil"
)

//address string, port int, username string, password string, destination string, soTimeOut int32, idleTimeOut int32

/*

// Config represents an Elasticsearch configuration.
type Config struct {
	URL         string
	Index       string
	Username    string
	Password    string
	Shards      int
	Replicas    int
	Sniff       *bool
	Healthcheck *bool
	Infolog     string
	Errorlog    string
	Tracelog    string
}

*/

type Conn struct {
	Value string `json:"value"`
}

func (l Conn) GetEsConfig() *esc.Config {
	c, _ := esc.Parse(l.Value)
	return c
}
