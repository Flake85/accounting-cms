package request

import "github.com/google/uuid"

type InvoiceRequest struct {
	ClientId   uuid.UUID `json:"clientId"`	
	IsInvoiced bool 	 `json:"IsInvoiced"`
}