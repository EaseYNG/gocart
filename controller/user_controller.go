package controller

import (
	"main/common"
	"main/dto"
	"main/service"

	"github.com/gin-gonic/gin"
)

// Register 注册
func Register(c *gin.Context) {
	var req dto.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Fail(c, err.Error())
		return
	}
	// 参数校验
	if err := req.Validate(); err != nil {
		common.Fail(c, err.Error())
		return
	}
	if bizErr := service.CreateUser(req); bizErr != nil {
		common.Fail(c, bizErr.Error())
		return
	}
	common.Success(c, "用户注册成功", nil)
}

// Login 登录
func Login(c *gin.Context) {
	var req dto.UserLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Fail(c, "参数错误")
		return
	}

	token, bizErr := service.LoginUser(req)
	if bizErr != nil {
		common.Fail(c, bizErr.Error())
		return
	}

	common.Success(c, "登录成功", gin.H{"token": token})
}

// Logout 注销
func Logout(c *gin.Context) {
	// 由于JWT是无状态的，实际项目中可以结合Redis黑名单实现真注销。
	// 这里通过响应让客户端主动清除Token即可
	common.Success(c, "注销成功", nil)
}
