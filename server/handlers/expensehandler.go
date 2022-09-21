package handlers

import (
	"encoding/json"
	"net/http"
	"server/model"
	"server/request"
	"server/response"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (handler *Handler) GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := handler.repository.GetAllExpenses()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving expenses", w)
		return
	}
	response.NewOkResponse(&expenses, w)
}

func (handler *Handler) GetExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	expense, err := handler.repository.FindExpenseByID(expenseId)
	if err != nil {
		response.NewErrorResponse(404, "expense not found", w)
		return
	}
	response.NewOkResponse(&expense, w)
}

func (handler *Handler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expenseReq request.ExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		response.NewErrorResponse(400, "expense decode malfunction", w)
		return
	}
	expense, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		response.NewErrorResponse(422, "expense validation error", w)
		return
	}
	expenseId, err := handler.repository.CreateExpense(&expense); if err != nil {
		response.NewErrorResponse(500, "error occurred creating expense", w)
		return
	}
	expense.ID = expenseId
	response.NewOkResponse(&expense, w)
}

func (handler *Handler) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	expense, err := handler.repository.FindExpenseByID(expenseId)
	if err != nil {
		response.NewErrorResponse(404, "expense not found", w)
		return
	}
	expenseReq := request.ExpenseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		response.NewErrorResponse(400, "error occurred decoding expense", w)
		return
	}
	expenseValidated, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		response.NewErrorResponse(422, "expense validation error", w)
		return
	}
	expense.Description = expenseValidated.Description
	expense.Cost = expenseValidated.Cost
	if err := handler.repository.UpdateExpense(&expense); err != nil {
		response.NewErrorResponse(500, "error occurred updating expense", w)
		return
	}
	response.NewOkResponse(&expense, w)
}

func (handler *Handler) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	_, err = handler.repository.FindExpenseByID(expenseId)
	if err != nil {
		response.NewErrorResponse(404, "expense not found", w)
		return
	}
	query := model.Expense{}
	query.ID = expenseId
	err = handler.repository.DeleteExpense(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
