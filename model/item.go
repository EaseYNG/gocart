package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string
	Price       float32
	Stock       int
	Description string
}
