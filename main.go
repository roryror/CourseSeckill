package main

import (
	route "course_seckill_clean_architecture/api/route"
	"course_seckill_clean_architecture/bootstrap"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Initializing application...")
	app := bootstrap.App()
	defer app.CloseConnections()

	env := app.Env
	
	db := app.MySQL
	cache := app.Redis
	mq := app.Kafka
	channel := app.Channel

	// 完全禁用Gin的日志输出
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
	log.SetOutput(io.Discard)
	
	router := gin.New() // 使用gin.New()而不是gin.Default()，避免默认的日志中间件
	router.Use(gin.Recovery()) // 仅添加恢复中间件，不添加日志中间件
	
	route.Setup(db, cache, mq, channel, router)
	route.Run(env.ServerPort, router)
}
