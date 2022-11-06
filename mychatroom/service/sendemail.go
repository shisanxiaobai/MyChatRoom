package service

import (
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils/errmsg"
)

// 检查邮箱
func CheckEmail(email string) int {

	var count int64
	if email == "" {
		return errmsg.ERROR_EMAIL_NOT_EMPTY
	}
	err := dao.Db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		log.Println("检查邮箱发生错误", err)
		return errmsg.ERROR_MYSQL
	} else if count > 0 {
		log.Printf("邮箱已存在")
		return errmsg.ERROR_EMAIL_EXIST
	}
	return errmsg.SUCCESS

}
