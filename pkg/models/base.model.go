package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IModel interface {
	Create()
	FindAll()
	FindOne()
	Update()
}

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {

	uid := uuid.NewV4()

	b.ID = uid

	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()

	return
}
