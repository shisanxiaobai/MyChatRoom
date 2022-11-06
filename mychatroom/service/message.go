package service

import (
	"context"
	"log"
	"mychatroom/dao"
	"mychatroom/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 向messagebasic集合里插入一条数据
func InsertMessage(message *model.Message) error {

	_, err := dao.Mdb.Collection(model.Message{}.CollectionName()).InsertOne(context.Background(), message)
	return err
}

func GetMessageListByRid(number string, limit, skip *int64) ([]model.Message, error) {
	var message []model.Message
	cursor, err := dao.Mdb.Collection(model.Message{}.CollectionName()).
		Find(context.Background(), bson.M{"number": number},
			&options.FindOptions{
				Limit: limit,
				Skip:  skip,
				Sort: bson.D{{
					Key: "creatde_at", Value: -1,
				}},
			})
	if err != nil {
		log.Println("查找信息列表失败", err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var msg model.Message
		err = cursor.Decode(msg)
		if err != nil {
			log.Println("message游标解码失败", err)
			return nil, err
		}
		message = append(message, msg)
	}

	return message, nil

}
