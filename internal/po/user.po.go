package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); not null; index:idx_uuid; unique;"`
	Username string    `gorm:"column:username; type:varchar(255)"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

func (user *User) TableName() string {
	return "go_crm_user"
}
