package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type Clicks struct {
	gorm.Model
	Ip      string
	ShortId int
}
