package service

import (
	"context"
	"log"
	"mychatroom/dao"
	"mychatroom/model"

	"go.mongodb.org/mongo-driver/bson"
)

// 添加一条UserRoom记录保存聊天室关联
func InsertOneUserRoom(ur *model.UserRoom) error {
	_, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).InsertOne(context.Background(), ur)
	return err
}

// 删除UserRoom关联记录
func DeleteUserRoom(number string) error {
	_, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		DeleteOne(context.Background(), bson.M{"number": number})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return err
	}
	return nil
}

// 根据账号和房间号查询UserRoom关联表
func GetUserRoomByAccountNumber(account, number string) (model.UserRoom, error) {
	var userRoom model.UserRoom
	err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "account", Value: account}, {Key: "number", Value: number}}).
		Decode(&userRoom)
	return userRoom, err
}

// 根据房间号查询UserRoom关联表
func GetUserRoomByNumber(number string) ([]model.UserRoom, error) {
	cursor, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{Key: "number", Value: number}})
	if err != nil {
		log.Println("查找UserRoom失败", err)
		return nil, err
	}
	var userRoom []model.UserRoom
	for cursor.Next(context.Background()) {
		var ur model.UserRoom
		err := cursor.Decode(&ur)
		if err != nil {
			log.Println("userRoom游标解码失败", err)
			return nil, err
		}
		userRoom = append(userRoom, ur)
	}
	return userRoom, nil
}

// 根据账号查找关联
func GetUserRoomAccount(account1, account2 string, roomType int) string {
	// 查询 account1 房间列表
	cursor, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{Key: "account", Value: account1}, {Key: "room_type", Value: roomType}})
	rooms := make([]string, 0)
	if err != nil {
		log.Println("[mongoDB ERROR]", err)
		return ""
	}

	for cursor.Next(context.Background()) {
		var ur *model.UserRoom
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return ""
		}
		rooms = append(rooms, ur.Number)
	}

	// 获取关联 account2 房间个数
	var ur *model.UserRoom
	err = dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.M{"account": account2, "room_type": roomType, "number": bson.M{"$in": rooms}}).Decode(ur)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return ""
	}

	return ur.Number
}

// 判断用户是否为好友(是否在同一私聊聊天室)
func JudgeUserIsFriend(account1, account2 string) bool {
	// 查询 account1 单聊房间列表 利用UserRoom关联表查询
	cursor, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{Key: "account", Value: account1}, {Key: "room_type", Value: 1}})
	roomNumbers := make([]string, 0)
	if err != nil {
		log.Printf("[MomgoDB ERROR]:%v\n", err)
		return false
	}
	for cursor.Next(context.Background()) {
		var ur *model.UserRoom
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return false
		}
		roomNumbers = append(roomNumbers, ur.Number)
	}
	// 获取关联 account2 单聊房间个数
	cnt, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"account": account2, "room_type": 1, "number": bson.M{"$in": roomNumbers}})
	if err != nil {
		log.Printf("[MongoDB ERROR]:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}

	return false
}

// 判断用户是否在群聊中
func IsUserJoin(account1, account2 string) bool {
	// 查询 account1 群聊房间列表 利用UserRoom关联表查询
	cursor, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{Key: "account", Value: account1}, {Key: "room_type", Value: 2}})
	roomNumbers := make([]string, 0)
	if err != nil {
		log.Printf("[MomgoDB ERROR]:%v\n", err)
		return false
	}
	for cursor.Next(context.Background()) {
		var ur *model.UserRoom
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return false
		}
		roomNumbers = append(roomNumbers, ur.Number)
	}
	// 获取关联 account2 群聊房间个数
	cnt, err := dao.Mdb.Collection(model.UserRoom{}.CollectionName()).
		CountDocuments(context.Background(), bson.M{"account": account2, "room_type": 2, "number": bson.M{"$in": roomNumbers}})
	if err != nil {
		log.Printf("[MongoDB ERROR]:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}

	return false
}
