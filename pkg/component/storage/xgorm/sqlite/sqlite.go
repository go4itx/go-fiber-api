package sqlite

import (
	"errors"
	conf "home/config"
	"home/pkg/component/storage/xgorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Build ...
func Build(name string) (db *gorm.DB, err error) {
	var config xgorm.Config
	if err = conf.Load(name, &config); err != nil {
		return
	}

	if config.Dsn == "" {
		err = errors.New("sqlite conf dns is empty")
		return
	}

	return xgorm.Init(sqlite.Open(config.Dsn), config)
}
