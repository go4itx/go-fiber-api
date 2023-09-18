package xgorm

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	Debug           bool
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int64
	TablePrefix     string
}

// Init ...
func Init(dialector gorm.Dialector, config Config) (db *gorm.DB, err error) {
	Logger := logger.Default
	if config.Debug {
		Logger = Logger.LogMode(logger.Info)
	}

	db, err = gorm.Open(dialector, &gorm.Config{
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
