package model

type Expense struct {
	BaseModel
	Description string
	Cost   		float64     
}