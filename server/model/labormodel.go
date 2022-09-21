package model

import "github.com/google/uuid"

type Labor struct {
	BaseModel
	Description string	   `json:"description"`
	ClientID	uuid.UUID  `json:"clientId" gorm:"TYPE:uuid"`
	Client		Client     `json:"client,omitempty" gorm:"constraint:OnDelete:CASCADE"`
	InvoiceID   *uuid.UUID `json:"invoiceId,omitempty" gorm:"TYPE:uuid"`
	Invoice		*Invoice   `json:"invoice,omitempty" gorm:"constraint:OnDelete:SET NULL"`
	HoursWorked float64	   `json:"hoursWorked"`
	HourlyRate  float64	   `json:"hourlyRate"`
	Total		float64	   `json:"total"`
}
