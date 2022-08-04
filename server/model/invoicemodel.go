package model

import (
	"github.com/google/uuid"
)

type Invoice struct {
	BaseModel
	ClientId   uuid.UUID
	IsInvoiced bool
}