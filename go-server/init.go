package main

import (
	"go_server/common/conf"
	"go_server/common/httpserv"
	"go_server/common/mysql"
)

func InitHandler() {
	conf.InitConfig()
	mysql.InitMysql()
	httpserv.InitHttpServer()
}
