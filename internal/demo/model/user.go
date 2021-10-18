package model

import (
	"gorm.io/gorm"
	"home/pkg/utils/password"
)

// User object
type User struct {
	gorm.Model
	RoleID   uint   `gorm:"type:int(10); unsigned; not null; comment:角色" json:"roleID"`
	Name     string `gorm:"type:varchar(20); unique; not null; comment:名称" json:"name"`
	Password string `gorm:"type:varchar(80);not null;comment:密码" json:"-"` //密码不返回
	Status   uint8  `gorm:"type:tinyint(1) unsigned NOT NULL;default:1;comment:状态：#1启用 #2禁用" json:"status"`
}

// BeforeCreate hook
func (m *User) BeforeCreate(db *gorm.DB) (err error) {
	if pwd, err := password.Hash(m.Password); err == nil {
		db.Statement.SetColumn("password", pwd)
	}

	return
}

// BeforeUpdate hook
func (m *User) BeforeUpdate(db *gorm.DB) (err error) {
	if m.Password != "" {
		if pwd, err := password.Hash(m.Password); err == nil {
			db.Statement.SetColumn("password", pwd)
		}
	}

	return
}
