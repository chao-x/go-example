package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_server/common/conf"
	"strconv"
)

var DB *gorm.DB

func InitMysql() {
	var err error

	viper := conf.Viper
	mysqlConf := viper.Get("mysql")
	mysqlConfMap := mysqlConf.(map[string]interface{})
	host := mysqlConfMap["host"].(string)
	port := mysqlConfMap["port"].(int64)
	portStr := strconv.FormatInt(port, 10)
	username := mysqlConfMap["username"].(string)
	password := mysqlConfMap["password"].(string)

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/kv?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, portStr)
	if err != nil {
		panic(err)
	}
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//DB.LogMode(true)
}
