// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGoCrmUser = "go_crm_user"

// GoCrmUser mapped from table <go_crm_user>
type GoCrmUser struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UUID      string         `gorm:"column:uuid;not null" json:"uuid"`
	Username  string         `gorm:"column:username" json:"username"`
	IsActive  bool           `gorm:"column:is_active" json:"is_active"`
}

// TableName GoCrmUser's table name
func (*GoCrmUser) TableName() string {
	return TableNameGoCrmUser
}
