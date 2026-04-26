package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: 200,
		Msg:  msg,
		Data: data,
	})
}

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: 400,
		Msg:  msg,
		Data: nil,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusOK, Result{
		Code: 401,
		Msg:  "未授权",
		Data: nil,
	})
}
