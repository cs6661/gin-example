package apis

import (
	"context"
	"gin-example/app/admin/models"
	"gin-example/app/admin/service"
	"gin-example/common/db"
	"gin-example/common/dto"
	"gin-example/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 放入底层初始化

type SysUserApi struct {
	Context context.Context
	Logger  *zap.Logger
	Orm     *gorm.DB
	Errors  error
}

func NewSysUserApi() *SysUserApi {
	return &SysUserApi{
		Context: context.Background(),
		Logger:  logger.Logger,
		Orm:     db.GormDB,
	}
}

var serv = service.NewSysUser()

func (s *SysUserApi) Get(c *gin.Context) {
	req := &models.SysUser{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		logger.Ctx(c).Error("参数错误", zap.Error(err))
		dto.ResponseError(c, dto.ParameterInvalid)
		return
	}
	res, err := serv.QueryUser(c, req)
	logger.Ctx(c).Info("fff")
	if err != nil {
		logger.Logger.Error("查询失败", zap.Error(err))
		dto.HttpError(c, dto.MySQLQueryError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
