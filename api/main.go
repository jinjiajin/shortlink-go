package api

import (
	"encoding/json"
	"fmt"
	"singo/conf"
	"singo/serializer"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
