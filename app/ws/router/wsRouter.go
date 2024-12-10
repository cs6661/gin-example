package router

import "github.com/gin-gonic/gin"

func WsRouter(r *gin.RouterGroup) {
	r.GET("/ws", func(c *gin.Context) {

	})
}
