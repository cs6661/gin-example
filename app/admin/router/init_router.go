package router

import (
	"gin-example/config"
	"gin-example/pkg/jwt"
	"time"

	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	var r *gin.Engine

	claims := &jwt.JwtClaims{}
	if config.Conf.Mode == "dev" {
		expireTime := time.Now().Add(24 * time.Hour)
		claims.ExpiresAt = jwtv5.NewNumericDate(expireTime)
	}
	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	InitExamplesRouter(r, claims)
}
