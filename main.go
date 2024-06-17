package main

import (
	"gin-example/config"
	"gin-example/middleware"
	"gin-example/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// 初始化日志
	logger.InitLogger(config.Conf.LogConfig)

	r.Use(middleware.RequestLogMiddleware())
	r.Use(middleware.GinRecovery(true))
	r.Run(":80")
}
