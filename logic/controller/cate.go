package controller

import (
	"GoPass/logic/controller/response"
	"GoPass/logic/model"
	_ "encoding/json"
	"github.com/labstack/echo/v4"
	_ "io/ioutil"
	"net/http"
)

type CateController struct{}

var Cate = &CateController{}

func (*CateController) List(c echo.Context) error {
	query := model.Cate{}.Model()
	var results []model.Cate
	query.Where(" deleted_at is null").Scan(&results)
	return c.JSON(http.StatusOK, response.Success(results))
}
