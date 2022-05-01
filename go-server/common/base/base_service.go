package base

import (
	"github.com/jinzhu/gorm"
	"go_server/common/mysql"
)

type BaseService struct {
}

func (bc *BaseService) GetDB() *gorm.DB {
	return mysql.DB
}
