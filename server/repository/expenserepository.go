package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func FindExpenseByID(expenseId uuid.UUID) (model.Expense, error) {
	expense := model.Expense{}
	if DB.Where("id = ?", expenseId).Find(&expense).Error != nil {
		return expense, fmt.Errorf("cannot find expense by id: %w", DB.Error)
	}
	return expense, nil
}

func GetAllExpenses() ([]model.Expense, error) {
	expenses := make([]model.Expense, 0)
	if DB.Find(&expenses).Error != nil {
		return expenses, fmt.Errorf("could not get expenses from db: %w", DB.Error)
	}
	return expenses, nil
}

func CreateExpense(expense *model.Expense) (uuid.UUID, error) {
	if DB.Create(expense).Error != nil {
		return expense.ID, fmt.Errorf("cannot create expense: %w", DB.Error)
	}
	return expense.ID, nil
}

func UpdateExpense(expense *model.Expense) error {
	if DB.Save(expense).Error != nil {
		return fmt.Errorf("cannot update expense: %w", DB.Error)
	}
	return nil
}

func DeleteExpense(expense *model.Expense) error {
	if DB.Delete(expense).Error != nil {
		return fmt.Errorf("cannot delete expense: %w", DB.Error)
	}
	return nil
}