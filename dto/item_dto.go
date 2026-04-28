package dto

import "main/common"

type ItemDto struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
}

func Validate(dto ItemDto) *common.BizError {
	if dto.Price < 0 {
		return common.New(400, "价格不能为负值")
	}
	if dto.Stock < 0 {
		return common.New(400, "库存不能为负值")
	}
	return nil
}
