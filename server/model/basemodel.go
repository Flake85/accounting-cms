package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID			uuid.UUID	`gorm:"types:uuid;primary_key" json:"id"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt"`
	DeletedAt	*time.Time	`sql:"index" json:"deletedAt"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", &id)
}