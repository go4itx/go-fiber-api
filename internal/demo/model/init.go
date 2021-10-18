package model

import (
	"gorm.io/gorm"
)

// Init model
func Init(db *gorm.DB) (err error) {
	// AutoMigrate data
	if !db.Migrator().HasTable(&User{}) {
		err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
		if err != nil {
			return
		}

		db.Create(&User{
			Name:     "admin",
			Password: "admin",
			RoleID:   1,
		})
	}

	return nil
}

// ParamID common param id
type ParamID struct {
	ID uint `form:"id" validate:"required" desc:"ID"`
}
