// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameBlockMetum = "block_meta"

// BlockMetum mapped from table <block_meta>
type BlockMetum struct {
	ID             string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	AccountID      string         `gorm:"column:account_id;not null" json:"account_id"`
	BlockTxLimit   int32          `gorm:"column:block_tx_limit;not null" json:"block_tx_limit"`
	TransitionTxID string         `gorm:"column:transition_tx_id;not null" json:"transition_tx_id"`
	OpeningDate    time.Time      `gorm:"column:opening_date;not null" json:"opening_date"`
	ClosingDate    time.Time      `gorm:"column:closing_date" json:"closing_date"`
}

// TableName BlockMetum's table name
func (*BlockMetum) TableName() string {
	return TableNameBlockMetum
}
