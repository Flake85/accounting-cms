package model

import "github.com/google/uuid"

type Labor struct {
	BaseModel
	Description string
	InvoiceId   uuid.UUID
	HoursWorked float64
	HourlyRate  float64
}