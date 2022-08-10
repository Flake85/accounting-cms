package validation

import (
	"errors"
	"server/model"
	"server/request"
)

func ExpenseValidation(expenseReq *request.ExpenseRequest) (expense model.Expense, err error) {
	if len(expenseReq.Description) < 1 {
		return expense, errors.New("description is required")
	}
	if (expenseReq.Cost < 0.01) {
		return expense, errors.New("cost is required")
	}
	expense = model.Expense{
		Description: expenseReq.Description,
		Cost: expenseReq.Cost,
	}
	return expense, err
}