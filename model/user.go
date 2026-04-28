package model

import (
	"main/config"

	"gorm.io/gorm"
)

type UserRole string // 用户角色

const (
	Customer UserRole = "customer"
	Merchant UserRole = "merchant"
	Manager  UserRole = "manager"
)

// 用户结构体
type User struct {
	gorm.Model
	Username string
	Password string
	Nickname string
	Role     UserRole `gorm:"type:varchar(10);default:'active'"`
}

func (role *UserRole) Value() string {
	return string(*role)
}

func Of(value string) UserRole {
	return UserRole(value)
}

// CreateUser 创建用户并写入数据库
// user 要创建的用户
// return 可能的数据库错误
func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

// GetUserByUsername 根据用户名查询用户
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// RemoveUser 删除用户
// user 要删除的用户
// return 可能的数据库错误
func RemoveUser(user *User) error {
	return config.DB.Delete(user).Error
}
