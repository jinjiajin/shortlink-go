package service

import (
	"os"
	"singo/model"
	"singo/serializer"
)

// ShortToLongService
type ShortToLongService struct {
	// Code string `form:"code" json:"code" binding:"required,min=5,max=500"`
}

// GetShort 通过短链接获取长链接
func (service *ShortToLongService) GetShort(code string, ip string) *serializer.Response {
	var shortUrl model.ShortUrl
	err := model.DB.Where("code = ?", code).First(&shortUrl).Error

	if err != nil {
		return &serializer.Response{
			Data: os.Getenv("HOST"),
		}
	}
	// 表单验证
	if err := service.saveClicks(shortUrl.ID, ip); err != nil {
		return err
	}

	return &serializer.Response{
		Data: shortUrl.LongUrl,
	}
}

func (service *ShortToLongService) saveClicks(shortId uint, ip string) *serializer.Response {
	var clicks model.Clicks
	clicks.ShortId = int(shortId)
	clicks.Ip = string(ip)
	// 保存数据
	if err := model.DB.Create(&clicks).Error; err != nil {
		return &serializer.Response{
			Code: 40001,
			Msg:  "保存失败",
		}
	}
	return nil
}
