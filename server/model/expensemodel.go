package model

type Expense struct {
	BaseModel
	Description string  `json:"description"`
	Cost   		float64 `json:"cost"`    
}