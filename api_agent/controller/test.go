package controller

import (
	"microservice_learning/api_agent/base"
	"github.com/gin-gonic/gin"
	"microservice_learning/api_agent/model"
)

type TestController struct {
	base.BaseController
}

func (c *TestController) Test(g *gin.Context) {
	data := g.DefaultQuery("data", "")
	info, err := model.TestModel(data)
	if err != nil {
		c.ResponseFailureForFuncErr(g, err.Error())
		return
	}
	c.ResponseData(g, info)
	return
}
