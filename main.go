package main

import (
	"gin-example/app/admin/router"
	"gin-example/common/db"
	"gin-example/config"
	"gin-example/middleware"
	"gin-example/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//var AppRouters = make([]func(), 0)

//func init() {
//	AppRouters = append(AppRouters, router.InitRouter)
//}

func main() {
	r := gin.New()
	// 初始化日志
	logger.InitLogger(config.Conf.LogConfig)
	// 初始化数据库连接
	if err := db.InitMysql(); err != nil {
		logger.Logger.Error("InitMysql failed", zap.Error(err))
	}
	r.Use(middleware.RequestLogMiddleware(),
		middleware.GinRecovery(true))
	router.InitRouter(r)
	if err := r.Run(":80"); err != nil {
		logger.Logger.Fatal("run fatal", zap.Error(err))
	}
}
