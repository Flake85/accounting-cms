package model

import "github.com/google/uuid"

type Sale struct {
	BaseModel
	ClientId	uuid.UUID
	InvoiceId	*uuid.UUID
	Description string
	Units		int8
	UnitCost	float64
}