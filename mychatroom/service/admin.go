package service

import (
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils/errmsg"

	"github.com/spf13/viper"
)

// 管理员身份检查
func CheckAdmin(username, password string) int {
	adminName := viper.GetString("admin.username")
	adminPassword := viper.GetString("admin.password")
	if username != adminName {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if password != adminPassword {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	return errmsg.SUCCESS
}

// 根据id获得账号
func GetUserAccountById(id int) (string, int) {
	var data model.User
	err := dao.Db.Select("account").Where("id = ?", id).First(&data).Error
	if err != nil {
		log.Println("根据id获得账号发生错误", err)
		return data.Account, errmsg.ERROR
	}

	return data.Account, errmsg.SUCCESS
}

// 删除账号
func AdminDeleteUser(id int) int {
	var data model.User
	err := dao.Db.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		log.Println("删除账号操作发生错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除聊天室
func AdminDeleteRoom(id int) int {
	var data model.Room
	err := dao.Db.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		log.Println("管理员删除聊天室操作发生错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获得用户列表
func GetUserList() ([]model.User, int) {
	var data []model.User
	err := dao.Db.Find(&data).Error
	if err != nil {
		log.Println("查询用户列表失败", err)
		return nil, errmsg.ERROR
	}
	return data, errmsg.SUCCESS
}
