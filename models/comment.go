package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	BlogId   uint
	ParentId uint
	NickName string
	Content  string
}
