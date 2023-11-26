// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	EmailAddress string         `gorm:"column:email_address;not null" json:"email_address"`
	FullName     string         `gorm:"column:full_name;not null" json:"full_name"`
	Password     string         `gorm:"column:password;not null" json:"password"`
	PhoneNumber  string         `gorm:"column:phone_number" json:"phone_number"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
