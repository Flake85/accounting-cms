package handlers

import (
	"encoding/json"
	"net/http"
	"server/model"
	"server/repository"
	"server/request"
	"server/response"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := repository.GetAllExpenses()
	if err != nil {
		res := response.NewErrorResponse(500, "error occurred retrieving expenses")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&expenses)
	json.NewEncoder(w).Encode(res.Body)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		res := response.NewErrorResponse(400, "invalid uuid")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expense, err := repository.FindExpenseByID(expenseId)
	if err != nil {
		res := response.NewErrorResponse(404, "expense not found")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&expense)
	json.NewEncoder(w).Encode(res.Body)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expenseReq request.ExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		res := response.NewErrorResponse(400, "expense decode malfunction")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expense, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		res := response.NewErrorResponse(422, "expense validation error")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expenseId, err := repository.CreateExpense(&expense); if err != nil {
		res := response.NewErrorResponse(500, "error occurred creating expense")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expense.ID = expenseId
	res := response.NewOkResponse(&expense)
	json.NewEncoder(w).Encode(res.Body)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		res := response.NewErrorResponse(400, "invalid uuid")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expense, err := repository.FindExpenseByID(expenseId)
	if err != nil {
		res := response.NewErrorResponse(404, "expense not found")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expenseReq := request.ExpenseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		res := response.NewErrorResponse(400, "error occurred decoding expense")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expenseValidated, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		res := response.NewErrorResponse(422, "expense validation error")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	expense.Description = expenseValidated.Description
	expense.Cost = expenseValidated.Cost
	if err := repository.UpdateExpense(&expense); err != nil {
		res := response.NewErrorResponse(500, "error occurred updating expense")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&expense)
	json.NewEncoder(w).Encode(res.Body)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		res := response.NewErrorResponse(400, "invalid uuid")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	query := model.Expense{}
	query.ID = expenseId
	err = repository.DeleteExpense(&query); if err != nil {
		res := response.NewErrorResponse(500, "invalid uuid")
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&query)
	json.NewEncoder(w).Encode(res.Body)
}
