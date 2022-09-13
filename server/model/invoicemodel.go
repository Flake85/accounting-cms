package model

import (
	"github.com/google/uuid"
)

type Invoice struct {
	BaseModel
	Description string	  `json:"description"`
	ClientID    uuid.UUID `json:"clientId" gorm:"TYPE:uuid"`
	Client	    Client	  `json:"client" gorm:"constraint:OnDelete:CASCADE;"`
	Sales		*[]Sale   `json:"sales"`
	SalesTotal	float64   `json:"salesTotal"`
	Labors		*[]Labor  `json:"labors"`
	LaborsTotal float64   `json:"laborsTotal"`
	GrandTotal	float64	  `json:"grandTotal"`
	IsPaid  	bool	  `json:"isPaid"`
}