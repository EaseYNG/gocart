package model

import (
	"errors"

	"main/config"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string
	Price       float64
	Stock       int
	Description string
	OwnerId     int
}

// FindById 根据商品id获取商品
func FindById(id int) (*Item, error) {
	var res *Item
	result := config.DB.Where("id = ?", id).First(res)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 未找到记录
			return nil, result.Error
		}
		return res, result.Error
	}
	return res, nil
}

// FindByOwnerId 根据merchant_id获取商品列表
func FindByOwnerId(id int) ([]*Item, error) {
	var res []*Item
	result := config.DB.Where("owner_id = ?", id).Find(res)
	if result.Error != nil {
		return res, result.Error
	}
	return res, nil
}

func Add(item *Item) error {
	return config.DB.Create(item).Error
}

func (item *Item) Remove() error {
	return config.DB.Delete(item).Error
}

func (item *Item) Update(temp *Item) error {
	return config.DB.Model(&temp).Updates(temp).Error
}
