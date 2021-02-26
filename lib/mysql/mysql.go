package mysql

import (
	mysqlConfig "GoPass/config/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)

var mysqlDatabases sync.Map

func InitMysqlConnect() {

	for k, v := range mysqlConfig.Config.Conns {

		fmt.Println(v.GetDsn())

		ConnectMysql(v.GetDsn(), k)
	}
}

func ConnectMysql(config string, name string) *gorm.DB {
	db, _ := gorm.Open("mysql", config)

	err := db.DB().Ping()
	if err != nil {
		panic("failed to connect mysql:" + name)
	}

	db.DB().SetMaxIdleConns(1024)
	db.DB().SetMaxOpenConns(1024)
	db.DB().SetConnMaxLifetime(time.Minute * 10) //连接超时10分钟，数据库的wait_timeout最好设置为11分钟

	mysqlDatabases.Store(name, db)
	return db
}

func Mysql(name string) *gorm.DB {
	db, _ := mysqlDatabases.Load(name)
	return db.(*gorm.DB)
}

func DisconnectMysql() {
	mysqlDatabases.Range(func(key, value interface{}) bool {
		defer value.(*gorm.DB).Close()
		return true
	})
}
