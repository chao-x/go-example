package dao

import (
	"github.com/jinzhu/gorm"
	"go_server/modules/site/model"
)

type SiteDao struct {
	DB *gorm.DB
}

func(sd *SiteDao) AddSite(site *model.Site) {
	sd.DB.AutoMigrate(&model.Site{})
	sd.DB.Create(&site)
}