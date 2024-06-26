package apis

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysUser struct {
	Context *gin.Context
	Logger  *zap.Logger
	Orm     *gorm.DB
	Errors  error
}
