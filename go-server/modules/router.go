package modules

import (
	"github.com/gin-gonic/gin"
	"go_server/modules/site/controller"
)

func ConfigureModuleRouter(R *gin.Engine) {
	R.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello wuchaoxin",
		})
	})
	R.GET("/site/add", new(controller.SiteController).AddSite)
}
