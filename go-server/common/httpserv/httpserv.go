package httpserv

import (
	"github.com/gin-gonic/gin"
	"go_server/common/conf"
	"go_server/modules"
	"strconv"
)

var Server *gin.Engine

func InitHttpServer() {
	port := conf.Viper.GetInt64("server.port")
	portStr := strconv.FormatInt(port, 10)

	Server = gin.Default()
	modules.ConfigureModuleRouter(Server)
	Server.Run(":" + portStr)
}
