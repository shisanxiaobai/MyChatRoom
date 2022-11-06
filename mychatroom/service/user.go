package service

import (
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils"
	"mychatroom/utils/errmsg"
)

// 注销用户
func DeleteUser(account string) int {
	var data model.User
	err := dao.Db.Where("account = ?", account).Delete(&data).Error
	if err != nil {
		log.Println("注销用户操作发生错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 更改用户个人资料
func EditUserInfo(id int, data *model.User) int {
	var user model.User
	var maps = make(map[string]interface{})
	maps["nickname"] = data.Nickname
	maps["sex"] = data.Sex
	maps["avatar"] = data.Avatar
	err := dao.Db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		log.Println("更改个人资料发生错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改密码
func ChangePassword(account, password string) int {
	var user model.User
	md5str := utils.Md5Encrypt(password)
	err := dao.Db.Model(&user).Where("account = ?", account).Update("password", md5str).Error
	if err != nil {
		log.Println("修改密码错误", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 根据账号获得用户信息
func GetUserInfoByAccount(account string) (model.User, int) {
	var user model.User
	err := dao.Db.Where("account = ?", account).First(&user).Error
	if err != nil {
		log.Println("根据账号查询用户信息失败", err)
		return user, errmsg.ERROR
	}

	return user, errmsg.SUCCESS
}

// 检查自己的房间
func CheckOwnRoom(account string) int {
	var count int64
	err := dao.Db.Model(&model.Room{}).Where("account = ?", account).Count(&count).Error
	if err != nil {
		log.Println("检查聊天室发生错误", err)
		return errmsg.ERROR_MYSQL
	} else if count > 0 {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 检查用户是否存在
func IsAccountExist(account string) bool {
	var data model.User
	err := dao.Db.Where("account =  ?", account).First(&data).Error
	if err != nil {
		log.Println("mysql查询错误", err)
		return false
	}

	return true
}

// 检查该聊天室是否属于用户
func IsRoomBelongMine(account, number string) bool {
	var data []model.Room
	err := dao.Db.Select("number").Where("account =  ?", account).Find(&data).Error
	if err != nil {
		log.Println("mysql查询错误", err)
		return false
	}

	for _, v := range data {
		if number == v.Number {
			return true
		}
	}

	return false
}
