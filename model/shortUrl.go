package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type ShortUrl struct {
	gorm.Model
	Code    string
	LongUrl string
}
