package service

import (
	"log"
	"mychatroom/dao"
	"mychatroom/model"
	"mychatroom/utils/errmsg"
)

// 创建聊天室
func CreateRoom(account, number, name, info string) int {
	data := &model.Room{
		Account: account,
		Number:  number,
		Name:    name,
		Info:    info,
	}
	err := dao.Db.Create(&data).Error
	if err != nil {
		log.Println("创建聊天室失败", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除聊天室
func DeleteRoom(number string) int {
	var data model.Room
	err := dao.Db.Where("number = ?", number).Delete(&data).Error
	if err != nil {
		log.Println("删除私人聊天室失败", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 更改聊天室信息
func EditRoom(number string, data *model.Room) int {
	var room model.Room
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	maps["info"] = data.Info
	err := dao.Db.Model(&room).Where("number = ?", number).Updates(maps).Error
	if err != nil {
		log.Println("修改聊天室信息失败", err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查看自己的聊天室信息
func OwnRoom(account string) ([]model.Room, int) {

	var room []model.Room
	err := dao.Db.Select("number", "Name", "Info").Where("account = ?", account).Order("article.created_at DESC").Find(&room).Error
	if err != nil {
		log.Println("查看聊天室信息失败", err)
		return nil, errmsg.ERROR
	}
	return room, errmsg.SUCCESS
}

// 查看所有聊天室列表
func RoomList() ([]model.Room, int) {
	var room []model.Room
	err := dao.Db.Select("number", "Name", "Info", "account").Order("article.created_at DESC").Find(&room).Error
	if err != nil {
		log.Println("查看聊天室列表失败", err)
		return nil, errmsg.ERROR
	}
	return room, errmsg.SUCCESS
}

// 检查聊天室是否存在
func CheckRoom(number string) int {
	var count int64
	err := dao.Db.Model(&model.Room{}).Where("number = ?", number).Count(&count).Error
	if err != nil {
		log.Println("检查聊天室发生错误", err)
		return errmsg.ERROR_MYSQL
	} else if count > 0 {
		return errmsg.ERROR_ROOM_EXIST
	}
	return errmsg.SUCCESS
}
