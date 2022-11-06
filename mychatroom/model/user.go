package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `swaggerignore:"true"`
	Account    string `gorm:"not null; index" json:"account"`
	Password   string `gorm:"not null" json:"password"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Email      string `gorm:"not null" json:"email"`
	Avatar     string `json:"avatar"`
}

// 查询指定用户资料存储在这个结构体
type UserQuery struct {
	Account  string `json:"account"`
	Nickname string `json:"nickname"`
	Sex      int    `son:"sex"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"` // 是否是好友 【true-是，false-否】
}
