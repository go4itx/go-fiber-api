package service

import (
	"gorm.io/gorm"
	"home/internal/demo/model"
	"home/pkg/storage/orm"
)

var (
	db   *gorm.DB
	User = newUserService()
)

// Init service
func Init() (err error) {
	if db, err = orm.Build("mysql.test"); err != nil {
		return
	}

	if err = model.Init(db); err != nil {
		return
	}

	return nil
}
