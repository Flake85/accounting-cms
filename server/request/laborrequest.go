package request

import "github.com/google/uuid"

type LaborRequest struct {
	Description	string	   `json:"description"`
	ClientId	uuid.UUID  `json:"clientId"`
	// InvoiceId	*uuid.UUID `json:"invoiceId,omitempty"`
	HoursWorked float64	   `json:"hoursWorked"`
	HourlyRate	float64	   `json:"hourlyRate"`
}