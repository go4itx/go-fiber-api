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

		db.Create(&[]User{
			{
				Name:     "test001",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test002",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test003",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test004",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test005",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test006",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test007",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test008",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test009",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test010",
				Password: "123456",
				RoleID:   1,
			},
			{
				Name:     "test011",
				Password: "123456",
				RoleID:   1,
			},
		})
	}

	return nil
}

// ParamID common param id
type ParamID struct {
	ID uint `form:"id" validate:"required" desc:"ID"`
}
