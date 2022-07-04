package api

import (
	"net/http"
	"shortlink/service"

	"github.com/gin-gonic/gin"
)

// LongToShort 生成短链接
func LongToShort(c *gin.Context) {
	var service service.LongToShortService
	if err := c.ShouldBind(&service); err == nil {
		res := service.ShortKeep()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShortToLong 短链接转换为长链接
func ShortToLong(c *gin.Context) {
	var service service.ShortToLongService
	res := service.GetShort(c.Param("code"), c.ClientIP())
	c.Redirect(http.StatusMovedPermanently, res.Data.(string))
	c.JSON(200, res)
}

// CountLinks 链接总数
func CountLinks(c *gin.Context) {
	var service service.CountLinksService
	res := service.GetCountLinks()
	c.JSON(200, res)
}

// CountClicks 链接点击次数
func CountClicks(c *gin.Context) {
	var service service.CountClicksService
	res := service.GetCountClicks()
	c.JSON(200, res)
}
