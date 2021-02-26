package controller

import (
	esm "GoPass/es/model"
	_ "GoPass/lib/helper"
	"GoPass/logic/controller/response"
	"GoPass/logic/model"
	"GoPass/service"
	_ "fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticleController struct{}

var Article = &ArticleController{}

func (*ArticleController) List(c echo.Context) error {
	params := map[string]interface{}{}
	if c.QueryParam("title") != "" {
		params["title"] = c.QueryParam("title")
	}
	if c.QueryParam("cate_id") != "" {
		params["cate_id"] = c.QueryParam("cate_id")
	}
	page, _ := strconv.Atoi(c.QueryParam("page"))
	size := 10
	///return c.JSON(http.StatusOK, response.Success(model.Article{}.List(params, page, size)))
	return c.JSON(http.StatusOK, response.Success(esm.Article{}.Search(params, page, size)))
}

func (*ArticleController) Info(c echo.Context) error {
	//	results := model.Article{}.Info(c.Param("id"))
	//results.SetReadAmount()
	//	return c.JSON(http.StatusOK, response.Success(results))
	esr := esm.Article{}.Get(c.Param("id"))

	if esr.Id == 0 {
		return c.JSON(http.StatusOK, response.Success(nil))
	}

	if esr.DeletedAt != nil || esr.Status == 0 {
		return c.JSON(http.StatusOK, response.Success(nil))
	}

	model.Article{}.SetReadAmountBy(c.Param("id"))
	service.ReadCountQueue.Push(c.Param("id"), -1)
	return c.JSON(http.StatusOK, response.Success(esr))
}
