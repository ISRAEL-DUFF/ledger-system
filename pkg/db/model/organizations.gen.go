// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrganization = "organizations"

// Organization mapped from table <organizations>
type Organization struct {
	ID           string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name         string         `gorm:"column:name;not null" json:"name"`
	Address      string         `gorm:"column:address;not null" json:"address"`
	EmailAddress string         `gorm:"column:email_address;not null" json:"email_address"`
	PhoneNumber  string         `gorm:"column:phone_number;not null" json:"phone_number"`
	OwnerID      string         `gorm:"column:owner_id;not null" json:"owner_id"`
}

// TableName Organization's table name
func (*Organization) TableName() string {
	return TableNameOrganization
}
