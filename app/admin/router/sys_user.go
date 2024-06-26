package router

import (
	"gin-example/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

func registerSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.JwtClaims) {
	r := v1.Group("/sys-user").Use(authMiddleware.JwtMiddleware())
	{
		r.GET("")
	}

}
