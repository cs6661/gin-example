package service

import (
	"database/sql"
	"gin-example/app/admin/models"
	"gorm.io/gorm"
)

type SysUser struct {
	Orm   *gorm.DB
	DB    *sql.DB
	Error error
}

func (s SysUser) QueryUser(req *models.SysUser) error {
	return nil
}
