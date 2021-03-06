package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RestControllerInterface interface {
	Get(c *gin.Context)
	Show(c *gin.Context)
	Edit(c *gin.Context)
	Store(c *gin.Context)
	Delete(c *gin.Context)
}

type RestController struct{}

type returnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r RestController) Forbidden(c *gin.Context) {
	c.String(http.StatusNotFound, "页面不存在！")
}

func (r RestController) Success(c *gin.Context, msg string, data interface{}) {
	var result returnData
	result = returnData{1, msg, data}
	c.JSON(http.StatusOK, result)
}

func (r RestController) Error(c *gin.Context, msg string, data interface{}) {
	var result returnData
	result = returnData{0, msg, data}
	c.JSON(http.StatusOK, result)
}
