package model

import "github.com/google/uuid"

type Sale struct {
	BaseModel
	InvoiceId	uuid.UUID
	Units		int8
	UnitCost	float64
}