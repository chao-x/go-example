package service

import (
	"go_server/common/base"
	"go_server/modules/site/dao"
	"go_server/modules/site/model"
)

type SiteService struct {
	*base.BaseService
}

func (ss *SiteService) AddSite(site *model.Site) {
	ss.GetDB()
	siteDao := dao.SiteDao{DB: ss.GetDB()}
	siteDao.AddSite(site)
}
