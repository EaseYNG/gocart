package main

import (
	"fmt"

	"main/config"
	"main/router"
)

func main() {
	fmt.Println("初始化logger: ")
	config.InitLogger()
	fmt.Println("初始化logger成功")
	fmt.Println("初始化数据库连接：")
	config.InitDB()
	fmt.Println("初始化数据库成功")
	fmt.Println("初始化路由: ")
	r := router.Init()
	_ = r.Run(":8080")
	fmt.Println("初始化路由成功")
}
