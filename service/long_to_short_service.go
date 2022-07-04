package service

import (
	"shortlink/model"
	"shortlink/serializer"
	"shortlink/util"
)

// LongToShortService
type LongToShortService struct {
	LongUrl string `form:"longUrl" json:"longUrl" binding:"required,url,min=1,max=500"`
}

// ShortKeep 存储长链接
func (service *LongToShortService) ShortKeep() serializer.Response {
	var shortUrlS model.ShortUrl
	err := model.DB.Where("long_url = ?", service.LongUrl).First(&shortUrlS).Error
	if err == nil {
		return serializer.BuildShortUrlResponse(shortUrlS)
	}

	//生成长链接
	url := service.LongUrl
	cb := func(url, keyword string) bool {
		// todo 查db或缓存判断keyword是否重复
		return true
	}
	surl := util.Generator(util.CHARSET_ALPHANUMERIC, url, cb)
	shortUrl := model.ShortUrl{
		LongUrl: service.LongUrl,
	}
	shortUrl.Code = surl

	// 保存数据
	if err := model.DB.Create(&shortUrl).Error; err != nil {
		return serializer.ParamErr("保存失败", err)
	}

	return serializer.BuildShortUrlResponse(shortUrl)
}
