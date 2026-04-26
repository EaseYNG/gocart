package controller

import (
	"log/slog"

	"main/common"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	slog.Info("Hello World!")
	common.Success(c, "Hello World!", nil)
}
