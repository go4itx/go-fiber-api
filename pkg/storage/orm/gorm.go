package orm

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"home/pkg/utils/conf"
	"time"
)

type Config struct {
	Debug           bool
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int64
	TablePrefix     string
}

func Build(name string) (db *gorm.DB, err error) {
	var config Config
	if err = conf.Load(name, &config); err != nil {
		return
	}

	if config.Dsn == "" {
		err = errors.New("gorm conf dns is empty")
		return
	}

	Logger := logger.Default
	if config.Debug {
		Logger = Logger.LogMode(logger.Info)
	}

	db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{
		Logger: Logger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   config.TablePrefix,
		},
	})

	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	if config.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}

	if config.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(100)
	}

	if config.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	}

	return
}
