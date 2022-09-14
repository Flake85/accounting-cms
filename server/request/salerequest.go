package request

import "github.com/google/uuid"

type SaleRequest struct {
	ClientId    uuid.UUID  `json:"clientId"`	
	Description string 	   `json:"description"`
	Units	    float64	   `json:"units"`
	UnitCost    float64	   `json:"unitCost"`
}