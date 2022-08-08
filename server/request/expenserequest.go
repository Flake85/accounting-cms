package request

type ExpenseRequest struct {
	Description	string  `json:"description"`
	Cost 		float64 `json:"cost"`
}