package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func FindSaleByID(saleId uuid.UUID) (model.Sale, error) {
	sale := model.Sale{}
	if DB.Where("id = ?", saleId).Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&sale).Error != nil {
		return sale, fmt.Errorf("Cannot find Sale by id: %w", DB.Error)
	}
	return sale, nil
}

func GetAllSales() ([]model.Sale, error) {
	sales := make([]model.Sale, 0)
	if DB.Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&sales).Error != nil {
		return sales, fmt.Errorf("could not get Sales from db: %w", DB.Error)
	}
	return sales, nil
}

func GetNotInvoicedSalesCountByClientId(clientId uuid.UUID) int64 {
	sales := make([]model.Sale, 0)
	result := DB.Where("invoice_id IS NULL AND client_id = ?", clientId).Find(&sales)
	return result.RowsAffected
}

func GetSalesByInvoiceId(invoiceId uuid.UUID) ([]model.Sale, error) {
	sales := make([]model.Sale, 0)
	if DB.Where("invoice_id = ?", invoiceId).Preload("Client").Find(&sales).Error != nil {
		return sales, fmt.Errorf("could not retrieve sales from db: %w", DB.Error)
	}
	return sales, nil
}

func UpdateSalesByClientId(clientId uuid.UUID, invoiceId uuid.UUID) (model.Sale, error) {
	sale := model.Sale{InvoiceID: &invoiceId}
	if DB.Model(&sale).Where("client_id = ? AND invoice_id IS NULL", clientId).Select("invoice_id").Updates(sale).Error != nil {
		return sale, fmt.Errorf("could not update sales by client id: %w", DB.Error)
	}
	return sale, nil
}

func CreateSale(sale *model.Sale) (uuid.UUID, error) {
	if DB.Create(sale).Error != nil {
		return sale.ID, fmt.Errorf("cannot create Sale: %w", DB.Error)
	}
	return sale.ID, nil
}

func UpdateSale(sale *model.Sale) error {
	if DB.Model(&sale).Select("description", "client_id", "units", "unit_cost", "total").Updates(sale).Error != nil {
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