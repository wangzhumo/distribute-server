package dao

import (
	"gorm.io/gorm"
)

type AppInfoDao struct {
	engine *gorm.DB
}

func NewAppInfoDao(db *gorm.DB) *AppInfoDao {
	return &AppInfoDao{
		engine: db,
	}
}

// func (dao *AppInfoDao) create(data *models.Apkinfo) error {

// }

// func (dao *AppInfoDao) query(data *models.Apkinfo) error {

// }
