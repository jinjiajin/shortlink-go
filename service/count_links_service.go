package service

import (
	"singo/model"
	"singo/serializer"
)

// CountLinksService
type CountLinksService struct {
}

// ShortKeep 存储长链接
func (service *CountLinksService) GetCountLinks() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.ShortUrl{}).Count(&count)
	// log.Println(count)
	return &serializer.Response{
		Data: count,
	}
}
