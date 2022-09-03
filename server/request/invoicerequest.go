package request

import "github.com/google/uuid"

type InvoiceRequest struct {
	ClientId   uuid.UUID `json:"clientId"`	
	IsPaid 	   bool 	 `json:"IsPaid"`
}