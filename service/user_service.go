package service

import (
	"log/slog"

	"main/common"
	"main/dto"
	"main/model"
	"main/util"
)

// CreateUser 创建用户
// uc controller传入的dto
// return 可能的错误
func CreateUser(uc dto.UserCreate) *common.BizError {
	hash, encode_err := util.Encode(uc.Password)
	if encode_err != nil {
		// BCrpt编码错误
		return common.InternalError
	}

	// 检查用户名是否已存在
	if _, err := model.GetUserByUsername(uc.Username); err == nil {
		return common.New(400, "用户名已存在")
	}

	if db_err := model.CreateUser(&model.User{
		Username: uc.Username,
		Password: hash,
		Nickname: uc.Nickname,
	}); db_err != nil {
		// db错误
		return common.InternalError
	}
	slog.Info("创建用户成功: " + uc.Username)
	return nil
}

// LoginUser 用户登录
func LoginUser(ul dto.UserLogin) (string, *common.BizError) {
	user, err := model.GetUserByUsername(ul.Username)
	if err != nil {
		return "", common.New(400, "用户名或密码错误")
	}

	if !util.Decode(ul.Password, user.Password) {
		return "", common.New(400, "用户名或密码错误")
	}

	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", common.InternalError
	}

	slog.Info("用户登录成功: " + ul.Username)
	return token, nil
}
