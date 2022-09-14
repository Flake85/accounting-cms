package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deletedAt"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	base.ID = id
	return err
}
