package main

import (
	"gin-example/app/admin/router"
	"gin-example/common/db"
	"gin-example/config"
	"gin-example/middleware"
	"gin-example/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//var AppRouters = make([]func(), 0)

//func init() {
//	AppRouters = append(AppRouters, router.InitRouter)
//}

func main() {
	r := gin.New()
	// 初始化日志
	logger.InitLogger(config.Conf.LogConfig)
	p := middleware.NewPrometheus("gin-example")
	p.Use(r)
	// 初始化数据库连接
	if err := db.InitMysql(); err != nil {
		logger.Logger.Error("InitMysql failed", zap.Error(err))
	}
	r.Use(middleware.RequestLogMiddleware(),
		middleware.GinRecovery(true))
	router.InitRouter(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	if err := r.Run(":80"); err != nil {
		logger.Logger.Fatal("run fatal", zap.Error(err))
	}
}
