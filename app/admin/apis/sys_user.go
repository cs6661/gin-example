package apis

import (
	"gin-example/app/admin/models"
	"gin-example/app/admin/service"
	"gin-example/common/dto"
	"gin-example/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 放入底层初始化

type SysUserApi struct {
	Context *gin.Context
	Logger  *zap.Logger
	Orm     *gorm.DB
	Errors  error
}

var serv = service.SysUser{}

func (s *SysUserApi) Get(c *gin.Context) {
	req := &models.SysUser{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		logger.Logger.Error("参数错误", zap.Error(err))
		dto.ResponseError(c, dto.ParameterInvalid)
		return
	}
	err = serv.QueryUser(req)
	if err != nil {
		logger.Logger.Error("查询失败", zap.Error(err))
		dto.HttpError(c, dto.MySQLQueryError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "")
}
