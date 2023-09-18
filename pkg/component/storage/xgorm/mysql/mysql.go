package mysql

import (
	"errors"
	"home/pkg/component/storage/xgorm"
	"home/pkg/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Build ...
func Build(name string) (db *gorm.DB, err error) {
	var config xgorm.Config
	if err = conf.Load(name, &config); err != nil {
		return
	}

	if config.Dsn == "" {
		err = errors.New("mysql conf dns is empty")
		return
	}

	return xgorm.Init(mysql.Open(config.Dsn), config)
}
