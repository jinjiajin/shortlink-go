package serializer

import (
	"os"
	"shortlink/model"
)

// ShortUrl 用户序列化器
type ShortUrl struct {
	Code string `json:"code"`
}

// BuildShortUrl 序列化用户
func BuildShortUrl(shortUrl model.ShortUrl) ShortUrl {
	return ShortUrl{
		Code: os.Getenv("HOST") + shortUrl.Code,
	}
}

// BuildShortUrlResponse 序列化用户响应
func BuildShortUrlResponse(shortUrl model.ShortUrl) Response {
	return Response{
		Data: BuildShortUrl(shortUrl),
	}
}
