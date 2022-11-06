package service

import (
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils"
	"mychatroom/utils/errmsg"
)

// 登录检查
func CheckLogin(account, password string) int {
	if account == "" || password == "" {
		return errmsg.ERROR_NOT_EMPTY
	} else if code := CheckAccount(account); code == errmsg.ERROR_USER_NOT_EXIST {
		return code
	} else if _, code = GetUserInfoByLogin(account, password); code == errmsg.ERROR_PASSWORD_WRONG {
		return code
	} else {
		return errmsg.SUCCESS
	}
}

// 判断密码是否正确
func GetUserInfoByLogin(account, password string) (model.User, int) {
	s := utils.Md5Encrypt(password)
	password = s
	var user model.User
	err := dao.Db.Where("account = ? AND password = ?", account, password).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}
