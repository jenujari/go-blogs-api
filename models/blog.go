package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	NickName string
	Title    string
	Content  string
}
