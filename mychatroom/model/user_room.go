package model

type UserRoom struct {
	Account   string `bson:"account"`
	Number    string `bson:"number"`
	RoomType  int    `bson:"room_type"` // 房间 类型 【1-独聊房间 2-群聊房间】
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}


