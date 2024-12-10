package db

import (
	"database/sql"
	"fmt"
	"gin-example/config"
	"gin-example/pkg/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *sql.DB
var GormDB *gorm.DB

func InitMysql() error {
	var err error
	mysqlConf := config.Conf.Mysql
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true",
		mysqlConf.User, mysqlConf.Password, mysqlConf.Address, mysqlConf.Name)
	DB, err = sql.Open("mysql", connect)
	if err != nil {
		logger.Logger.Error("mysql connect err", zap.Error(err))
		return err
	}
	newLogger := NewLogger(Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  Info,
		IgnoreRecordNotFoundError: false,
	})
	if GormDB, err = gorm.Open(mysql.Open(connect), &gorm.Config{Logger: newLogger}); err != nil {
		return err
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return nil
}
