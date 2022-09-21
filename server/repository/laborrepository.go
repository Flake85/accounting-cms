package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func (repo *Repository) FindLaborByID(laborId uuid.UUID) (model.Labor, error) {
	labor := model.Labor{}
	if repo.db.Where("id = ?", laborId).Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&labor).Error != nil {
		return labor, fmt.Errorf("Cannot find labor by id: %w", repo.db.Error)
	}
	return labor, nil
}

func (repo *Repository) GetAllLabors() ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	if repo.db.Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&labors).Error != nil {
		return labors, fmt.Errorf("could not get labors from repo.db: %w", repo.db.Error)
	}
	return labors, nil
}

func (repo *Repository) GetNotInvoicedLaborsCountByClientId(clientId uuid.UUID) int64 {
	labors := make([]model.Labor, 0)
	result := repo.db.Where("invoice_id IS NULL AND client_id = ?", clientId).Find(&labors)
	return result.RowsAffected
}

func (repo *Repository) GetLaborsByInvoiceId(invoiceId uuid.UUID) ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	if repo.db.Where("invoice_id = ?", invoiceId).Preload("Client").Find(&labors).Error != nil {
		return labors, fmt.Errorf("could not retrieve labors from repo.db: %w", repo.db.Error)
	}
	return labors, nil
}

func (repo *Repository) UpdateLaborsByClientId(clientId uuid.UUID, invoiceId uuid.UUID) (model.Labor, error) {
	labor := model.Labor{InvoiceID: &invoiceId}
	if repo.db.Model(&labor).Where("client_id = ? AND invoice_id IS NULL", clientId).Select("invoice_id").Updates(labor).Error != nil {
		return labor, fmt.Errorf("could not update labors by client id: %w", repo.db.Error)
	}
	return labor, nil
}

func (repo *Repository) CreateLabor(labor *model.Labor) (uuid.UUID, error) {
	if repo.db.Create(labor).Error != nil {
		return labor.ID, fmt.Errorf("cannot create labor: %w", repo.db.Error)
	}
	return labor.ID, nil
}

func (repo *Repository) UpdateLabor(labor *model.Labor) error {
	if repo.db.Model(&labor).Select("description", "client_id", "hours_worked", "hourly_rate", "total").Updates(labor).Error != nil {
		return fmt.Errorf("cannot update labor: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) DeleteLabor(labor *model.Labor) error {
	if repo.db.Delete(labor).Error != nil {
		return fmt.Errorf("cannot delete labor: %w", repo.db.Error)
	}
	return nil
}
