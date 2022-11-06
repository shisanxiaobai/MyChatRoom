package model

type Message struct {
	Account   string `bson:"account"`
	Number    string `bson:"number"`
	Data      string `bson:"data"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (Message) CollectionName() string {
	return "message"
}

type InstantMessage struct {
	Number  string `json:"number"`
	Message string `json:"message"`
}
