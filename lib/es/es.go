package es

import (
	esConfig "GoPass/config/es"
	"context"
	_ "encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"sync"
)

var esDatabases sync.Map

func InitEsConnect() {
	for k, v := range esConfig.Config.Conns {
		fmt.Println(v.GetEsConfig())
		ConnectEs(v.GetEsConfig(), k)
	}
}

func ConnectEs(cc *config.Config, name string) *elastic.Client {

	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err := elastic.NewClientFromConfig(cc)
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(cc.URL).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	_, err = client.ElasticsearchVersion(cc.URL)
	if err != nil {
		panic(err)
	}
	esDatabases.Store(name, client)
	return client
}

func Es(name string) *elastic.Client {
	client, _ := esDatabases.Load(name)
	return client.(*elastic.Client)
}

func DisconnectEs() {
	esDatabases.Range(func(key, value interface{}) bool {
		defer value.(*elastic.Client).Stop()
		return true
	})
}
