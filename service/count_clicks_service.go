package service

import (
	"shortlink/model"
	"shortlink/serializer"
)

// CountClicksService
type CountClicksService struct {
}

// ShortKeep 存储长链接
func (service *CountClicksService) GetCountClicks() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.Clicks{}).Count(&count)
	// log.Println(count)
	return &serializer.Response{
		Data: count,
	}
}
