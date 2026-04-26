package router

import (
	"main/controller"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	public := r.Group("api")
	{
		public.GET("/test", controller.Hello)
		public.POST("/user/register", controller.Register)
		public.POST("/user/login", controller.Login)
	}

	// 需要鉴权的接口
	auth := r.Group("api")
	auth.Use(middleware.JWTAuth())
	{
		auth.POST("/user/logout", controller.Logout)
	}

	return r
}
