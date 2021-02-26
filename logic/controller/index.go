package controller

import (
	"GoPass/logic/controller/response"
	"GoPass/logic/model"
	_ "encoding/json"
	"github.com/labstack/echo/v4"
	_ "io/ioutil"
	"net/http"
)

const indexPics = "IndexPics"
const websiteInfo = "WebsiteInfo"

type indexController struct{}

var Index = &indexController{}

func (*indexController) GetSetting(c echo.Context) error {
	setting := model.Setting{}
	setting.Model().Where("key_name = ?", c.Param("key")).First(&setting)
	return c.JSON(http.StatusOK, response.Success(setting.Content))
}

func (*indexController) GetFeatureSetting(c echo.Context) error {
	setting := model.Setting{}
	setting.Model().Where("key_name = ?", "Feature").First(&setting)
	var res []model.Article
	model.Article{}.Model().Where("id IN (?)", setting.Content.Get()).Where("deleted_at is null").Select("title,image,id").Find(&res)
	return c.JSON(http.StatusOK, response.Success(res))
}
