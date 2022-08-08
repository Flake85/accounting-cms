package model

import "github.com/google/uuid"

type Labor struct {
	BaseModel
	Description string
	ClientId	uuid.UUID
	InvoiceId   *uuid.UUID
	HoursWorked float64
	HourlyRate  float64
}