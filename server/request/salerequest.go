package request

import "github.com/google/uuid"

type SaleRequest struct {
	ClientId    uuid.UUID  `json:"clientId"`	
	InvoiceId   *uuid.UUID `json:"invoiceId,omitempty"`
	Description string 	   `json:"description"`
	Units	    float64	   `json:"units"`
	UnitCost    float64	   `json:"unitCost"`
}