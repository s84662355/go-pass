package main

import (
	"GoPass/config"
	_ "GoPass/config/nsq"
	"GoPass/config/redis"
	"GoPass/lib/es"
	"GoPass/lib/mysql"
	"GoPass/logic/controller"
	"GoPass/service"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	es.InitEsConnect()
	mysql.ConnectMysql(config.MySQL, "default")
	defer mysql.DisconnectMysql()

	service.GlobalInit()

	fmt.Println(redis.Config)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	eee := e.Group("/blog/api")
	eee.GET("/article", controller.Article.List)

	eee.GET("/setting/:key", controller.Index.GetSetting)

	eee.GET("/article/:id", controller.Article.Info)
	eee.GET("/cate", controller.Cate.List)
	eee.GET("/GetFeatureSetting", controller.Index.GetFeatureSetting)

	e.Logger.Fatal(e.Start(":1323"))
}
