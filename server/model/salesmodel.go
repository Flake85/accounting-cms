package model

import "github.com/google/uuid"

type Sale struct {
	BaseModel
	Description string	   `json:"description"`
	ClientID	uuid.UUID  `json:"clientId" gorm:"TYPE:uuid"`
	Client	    Client	   `json:"client" gorm:"constraint:OnDelete:CASCADE;"`
	InvoiceID	*uuid.UUID `json:"invoiceId,omitempty" gorm:"TYPE:uuid"`
	Invoice		*Invoice   `json:"invoice,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
	Units		float64	   `json:"units"`
	UnitCost	float64	   `json:"unitCost"`
	Total		float64	   `json:"total"`	
}
