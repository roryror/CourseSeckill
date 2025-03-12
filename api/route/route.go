package route

import (
	"context"
	ctrler "course_seckill_clean_architecture/api/controller"
	domain "course_seckill_clean_architecture/domain"
	interfaces "course_seckill_clean_architecture/interface"
	"net/http"

	"github.com/gin-gonic/gin"
)


var ctx context.Context
var internal *domain.Internals
var controller *ctrler.Controller

func Setup(db interfaces.Database, cache interfaces.Cache, mq interfaces.MsgQueue, channel interfaces.Channel, router *gin.Engine) {
	ctx = context.Background()
	internal = &domain.Internals{Db: db, Cache: cache, Mq: mq, Channel: channel}
	controller = ctrler.NewController(ctx, internal)
	
	// 设置静态文件服务
	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/index.html")
	})
	
	NewSeckillRoute(router)
	NewViewerRoute(router)
	NewWarmupRoute(router)
}

func Run(address string, router *gin.Engine) {
	router.Run(address)
}