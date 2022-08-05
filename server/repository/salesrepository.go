package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func FindSaleByID(saleId uuid.UUID) (model.Sale, error) {
	sale := model.Sale{}
	if DB.Where("id = ?", saleId).Find(&sale).Error != nil {
		return sale, fmt.Errorf("Cannot find Sale by id: %w", DB.Error)
	}
	return sale, nil
}

func GetAllSales() ([]model.Sale, error) {
	sales := make([]model.Sale, 0)
	if DB.Find(&sales).Error != nil {
		return sales, fmt.Errorf("could not get Sales from db: %w", DB.Error)
	}
	return sales, nil
}

func CreateSale(sale *model.Sale) (uuid.UUID, error) {
	if DB.Create(sale).Error != nil {
		return sale.ID, fmt.Errorf("cannot create Sale: %w", DB.Error)
	}
	return sale.ID, nil
}

func UpdateSale(sale *model.Sale) error {
	if DB.Save(sale).Error != nil {
		return fmt.Errorf("cannot update Sale: %w", DB.Error)
	}
	return nil
}

func DeleteSale(sale *model.Sale) error {
	if DB.Delete(sale).Error != nil {
		return fmt.Errorf("cannot delete Sale: %w", DB.Error)
	}
	return nil
}