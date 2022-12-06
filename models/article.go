package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	NickName string
	Title    string
	Content  string
}
