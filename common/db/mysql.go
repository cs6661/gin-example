package db

import (
	"database/sql"
	"fmt"
	"gin-example/config"
)

var DB *sql.DB

func initMysql() {
	var err error
	mysql := config.Conf.Mysql
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true",
		mysql.User, mysql.Password, mysql.Address, mysql.Name)
	DB, err = sql.Open("mysql", connect)
	if err != nil {

	}
}
