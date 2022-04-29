package controller

import (
	"github.com/gin-gonic/gin"
	"go_server/common/base"
	"go_server/modules/site/model"
	"net/http"
)

type SiteController struct {
	*base.BaseController
}

func (sc *SiteController) AddSite(c *gin.Context) {
	var params model.Site
	var err error
	err = c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest,&base.ErrUnmarshalJsonParam)
		return
	}
}

