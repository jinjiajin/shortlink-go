package server

import (
	"net/http"
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	//加载静态文件
	r.Static("/static", "view/static")
	//模板解析
	r.LoadHTMLGlob("view/index.html")
	//模板渲染
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// 获取地址
	r.GET("/:code", api.ShortToLong)
	// 生成短链接
	r.POST("/api/links", api.LongToShort)
	// 链接总数
	r.GET("/api/count/links", api.CountLinks)
	// 链接点击次数
	r.GET("/api/count/clicks", api.CountClicks)
	return r
}
