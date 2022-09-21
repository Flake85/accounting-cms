package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func (repo *Repository) FindExpenseByID(expenseId uuid.UUID) (model.Expense, error) {
	expense := model.Expense{}
	if repo.db.Where("id = ?", expenseId).Find(&expense).Error != nil {
		return expense, fmt.Errorf("cannot find expense by id: %w", repo.db.Error)
	}
	return expense, nil
}

func (repo *Repository) GetAllExpenses() ([]model.Expense, error) {
	expenses := make([]model.Expense, 0)
	if repo.db.Find(&expenses).Error != nil {
		return expenses, fmt.Errorf("could not get expenses from repo.db: %w", repo.db.Error)
	}
	return expenses, nil
}

func (repo *Repository) CreateExpense(expense *model.Expense) (uuid.UUID, error) {
	if repo.db.Create(expense).Error != nil {
		return expense.ID, fmt.Errorf("cannot create expense: %w", repo.db.Error)
	}
	return expense.ID, nil
}

func (repo *Repository) UpdateExpense(expense *model.Expense) error {
	if repo.db.Save(expense).Error != nil {
		return fmt.Errorf("cannot update expense: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) DeleteExpense(expense *model.Expense) error {
	if repo.db.Delete(expense).Error != nil {
		return fmt.Errorf("cannot delete expense: %w", repo.db.Error)
	}
	return nil
}
