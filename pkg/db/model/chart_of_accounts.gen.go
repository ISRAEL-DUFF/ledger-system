// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameChartOfAccount = "chart_of_accounts"

// ChartOfAccount mapped from table <chart_of_accounts>
type ChartOfAccount struct {
	ID            string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	AccountNumber string         `gorm:"column:account_number;not null" json:"account_number"`
	Description   string         `gorm:"column:description;not null" json:"description"`
	Type          string         `gorm:"column:type;not null" json:"type"`
}

// TableName ChartOfAccount's table name
func (*ChartOfAccount) TableName() string {
	return TableNameChartOfAccount
}
