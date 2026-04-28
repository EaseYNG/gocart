package config

// 通过go get gorm.io/gorm, go get gorm.io/driver/mysql安装依赖

import (
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	password := "20060221@yyx"
	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/gocart_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("数据库连接失败! " + err.Error())
		panic("数据库连接失败! " + err.Error())
	}

	DB = db
}
