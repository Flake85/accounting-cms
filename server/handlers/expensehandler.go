package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/model"
	"server/repository"
	"server/request"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	expenses, err := repository.GetAllExpenses()
	if err != nil {
		log.Println("error occurred getting expenses")
		w.WriteHeader(500)
		fmt.Fprintln(w, "error occurred retrieving expenses")
		return
	}
	json.NewEncoder(w).Encode(&expenses)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", expenseIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	expense, err := repository.FindExpenseByID(expenseId)
	if err != nil {
		log.Printf("expense: %v, not found", expenseIdParam)
		w.WriteHeader(404)
		fmt.Fprintln(w, "expense not found")
		return
	}
	json.NewEncoder(w).Encode(&expense)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expenseReq request.ExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		log.Print("Expense decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred creating expense")
		return
	}
	expense, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		log.Println("validation error. expense not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "client validation error: %v", err)
		return
	}
	expenseId, err := repository.CreateExpense(&expense); if err != nil {
		log.Printf("expense: %v, not created", expenseId)
		w.WriteHeader(500)
		fmt.Fprintf(w, "an error occurred creating expense: %v", err)
		return
	}
	expense.ID = expenseId
	json.NewEncoder(w).Encode(&expense)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", expenseIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating expense")
		return
	}
	expense, err := repository.FindExpenseByID(expenseId)
	if err != nil {
		log.Printf("expense not found with uuid: %v", expenseId)
		w.WriteHeader(404)
		fmt.Fprintln(w, "expense not found")
		return
	}
	expenseReq := request.ExpenseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&expenseReq); err != nil {
		log.Print("expense decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating expense")
		return
	}
	expenseValidated, err := validation.ExpenseValidation(&expenseReq); if err != nil {
		log.Println("validation error. expense not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "expense validation error: %v", err)
		return
	}
	expense.Description = expenseValidated.Description
	expense.Cost = expenseValidated.Cost
	if err := repository.UpdateExpense(&expense); err != nil {
		log.Printf("error occurred updating expense id: %v", expenseId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred updating expense")
		return
	}
	json.NewEncoder(w).Encode(&expense)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	expenseIdParam := mux.Vars(r)["id"]
	expenseId, err := uuid.Parse(expenseIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", expenseIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	query := model.Expense{}
	query.ID = expenseId
	err = repository.DeleteExpense(&query); if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	json.NewEncoder(w).Encode(&query)
}
