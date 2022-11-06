package service

import (
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils"
	"mychatroom/utils/errmsg"
)

// 检查账号
func CheckAccount(account string) int {
	var count int64
	err := dao.Db.Model(&model.User{}).Where("account = ?", account).Count(&count).Error
	if err != nil {
		log.Println("检查账号发生错误", err)
		return errmsg.ERROR_MYSQL
	} else if count > 0 {
		return errmsg.ERROR_ACCOUNT_USED
	}
	return errmsg.SUCCESS
}

// 注册账号
func Register(data *model.User) int {
	password := utils.Md5Encrypt(data.Password)
	data.Password = password
	err := dao.Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
