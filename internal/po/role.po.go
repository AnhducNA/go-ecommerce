package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id; type:int; not null; primaryKey; autoIncrement;"`
	RoleName string `gorm:"column:role_name; type:varchar(255);"`
	RoleNote string `gorm:"column:role_note; type:text;"`
}

func (role *Role) TableName() string {
	return "role"
}
