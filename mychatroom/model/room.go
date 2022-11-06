package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Number  string `gorm:"unqiue; not null; index" json:"number"`
	Name    string `json:"name"`
	Info    string `json:"info"`
	Account string `gorm:"not null; index" json:"account"`
}
