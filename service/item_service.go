package service

import (
	"log/slog"

	"main/common"
	"main/dto"
	"main/model"
)

func AddItem(d dto.ItemDto) *common.BizError {
	if err := dto.Validate(d); err != nil {
		return err
	}
	if err := model.Add(&model.Item{
		Name:        d.Name,
		Price:       d.Price,
		Stock:       d.Stock,
		Description: d.Description,
	}); err != nil {
		return common.InternalError // DB错误
	}
	slog.Info("添加商品")
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
