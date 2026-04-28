package controller

import (
	"main/common"
	"main/dto"
	"main/model"
	"main/service"

	"github.com/gin-gonic/gin"
)

func AddItem(c *gin.Context) {
	ownerId := c.GetInt("userID")
	role := c.GetString("userRole")
	if model.Of(role) != model.Merchant {
		common.Fail(c, "用户权限错误")
		return
	}
	var itemDto dto.ItemDto
	if c.ShouldBindJSON(&itemDto) != nil {
		common.Fail(c, "表单错误：检查缺失的字段")
		return
	}
	if dto.Validate(itemDto) != nil {
		common.Fail(c, "表单错误：检查price、stock字段应不为负值")
		return
	}
	if err := service.AddItem(ownerId, itemDto); err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, "添加商品成功", nil)
}

func GetItemListMerchant(c *gin.Context) {
	ownerId := c.GetInt("userID")
	var list []dto.ItemDto
	list, err := service.GetItemList(ownerId)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, "获取商品列表成功", list)
}

func GetItemListCustomer(c *gin.Context, ownerId int) {
	var list []dto.ItemDto
	list, err := service.GetItemList(ownerId)
	if err != nil {
		common.Fail(c, err.Error())
		return
	}
	common.Success(c, "获取商品列表成功", list)
}
