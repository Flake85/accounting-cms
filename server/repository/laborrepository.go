package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)


func FindLaborByID(laborId uuid.UUID) (model.Labor, error) {
	labor := model.Labor{}
	if DB.Where("id = ?", laborId).Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&labor).Error != nil {
		return labor, fmt.Errorf("Cannot find labor by id: %w", DB.Error)
	}
	return labor, nil
}

func GetAllLabors() ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	if DB.Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&labors).Error != nil {
		return labors, fmt.Errorf("could not get labors from db: %w", DB.Error)
	}
	return labors, nil
}

func GetLaborsByInvoiceId(invoiceId uuid.UUID) ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	if DB.Where("invoice_id = ?", invoiceId).Preload("Client").Find(&labors).Error != nil {
		return labors, fmt.Errorf("could not retrieve labors from db: %w", DB.Error)
	}
	return labors, nil
}

// TODO: returns an empty array but the update is functional
func UpdateLaborsByClientId(clientId uuid.UUID, invoiceId uuid.UUID) ([]model.Labor, error) {
	labors := make([]model.Labor, 0)
	labor := model.Labor{InvoiceID: &invoiceId}
	if DB.Model(labor).Where("client_id = ? AND invoice_id IS NULL", clientId).Select("invoice_id").Updates(labor).Error != nil {
		return labors, fmt.Errorf("could not update labors by client id: %w", DB.Error)
	}
	return labors, nil
}

func CreateLabor(labor *model.Labor) (uuid.UUID, error) {
	if DB.Create(labor).Error != nil {
		return labor.ID, fmt.Errorf("cannot create labor: %w", DB.Error)
	}
	return labor.ID, nil
}

func UpdateLabor(labor *model.Labor) error {
	if DB.Save(labor).Error != nil {
		return fmt.Errorf("cannot update labor: %w", DB.Error)
	}
	return nil
}

func DeleteLabor(labor *model.Labor) error {
	if DB.Delete(labor).Error != nil {
		return fmt.Errorf("cannot delete labor: %w", DB.Error)
	}
	return nil
}