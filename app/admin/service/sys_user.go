package service

import (
	"database/sql"
	"gin-example/app/admin/models"
	"gin-example/common/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SysUser struct {
	Orm   *gorm.DB
	DB    *sql.DB
	Error error
}

func NewSysUser() *SysUser {
	return &SysUser{
		Orm:   db.GormDB,
		DB:    db.DB,
		Error: nil,
	}
}

func (s *SysUser) QueryUser(c *gin.Context, req *models.SysUser) ([]models.User, error) {
	res := make([]models.User, 0)
	err := db.GormDB.WithContext(c.Request.Context()).Limit(100).Find(&res).Error
	//err := s.Orm.First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
