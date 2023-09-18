package model

import (
	"gorm.io/gorm"
)

var db *gorm.DB

// Init model
func Init(gormDB *gorm.DB) (err error) {
	db = gormDB
	// AutoMigrate data
	if !db.Migrator().HasTable(&User{}) {
		err = db.AutoMigrate(&User{})
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

// OrderBy common struct
type OrderBy struct {
	Field     string `json:"field" validate:"required"`
	Direction string `json:"direction" validate:"required,oneof=asc desc"`
}
