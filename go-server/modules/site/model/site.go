package model

import "github.com/jinzhu/gorm"

type Site struct {
	gorm.Model
	SiteId   int64
	SiteName string
	SiteStr  string
}
