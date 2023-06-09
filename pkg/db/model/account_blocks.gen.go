// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAccountBlock = "account_blocks"

// AccountBlock mapped from table <account_blocks>
type AccountBlock struct {
	ID                string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt         time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	AccountID         string         `gorm:"column:account_id;not null" json:"account_id"`
	IsCurrentBlock    bool           `gorm:"column:is_current_block;not null" json:"is_current_block"`
	BlockSize         int32          `gorm:"column:block_size;not null" json:"block_size"`
	TransactionsCount int32          `gorm:"column:transactions_count;not null" json:"transactions_count"`
	Status            string         `gorm:"column:status;not null" json:"status"`
}

// TableName AccountBlock's table name
func (*AccountBlock) TableName() string {
	return TableNameAccountBlock
}
