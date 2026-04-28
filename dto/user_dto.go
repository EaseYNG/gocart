package dto

import "errors"

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserCreate) Validate() error {
	if uc.Username == "" {
		return errors.New("用户名为空！")
	}
	if uc.Password == "" {
		return errors.New("密码为空！")
	}
	return nil
}
