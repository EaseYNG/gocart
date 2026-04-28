package service

import (
	"log/slog"

	"main/common"
	"main/dto"
	"main/model"
)

// AddItem 传入商家id和dto构建item对象加入到数据库中
// return 可能的业务错误
func AddItem(ownerId int, d dto.ItemDto) *common.BizError {
	if err := dto.Validate(d); err != nil {
		return err
	}
	if err := model.Add(&model.Item{
		Name:        d.Name,
		Price:       d.Price,
		Stock:       d.Stock,
		Description: d.Description,
		OwnerId:     ownerId,
	}); err != nil {
		return common.InternalError // DB错误
	}
	slog.Info("添加商品", "owner_id", ownerId)
	return nil
}

func UpdateItem(itemId int, d dto.ItemDto) *common.BizError {
	if err := dto.Validate(d); err != nil {
		return err
	}
	item, err := model.FindById(itemId)
	if err == nil {
		item.Update(&model.Item{
			Name:        d.Name,
			Price:       d.Price,
			Stock:       d.Stock,
			Description: d.Description,
		})
		slog.Info("更新商品", "item_id", itemId)
	} else {
		return common.InternalError // DB错误
	}
	return nil
}

// GetItemList 获取当前商家的全部商品（买家用户可使用）
// return 商品dto列表
func GetItemList(ownerId int) ([]dto.ItemDto, *common.BizError) {
	items, err := model.FindByOwnerId(ownerId)
	if err != nil {
		return nil, common.InternalError
	}
	itemDtos := make([]dto.ItemDto, len(items))
	for i := 0; i < len(items); i++ {
		itemDtos[i] = dto.ItemDto{
			Name:        items[i].Name,
			Price:       items[i].Price,
			Stock:       items[i].Stock,
			Description: items[i].Description,
		}
	}
	slog.Info("获取商品列表")
	return itemDtos, nil
}
