package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func (repo *Repository) FindSaleByID(saleId uuid.UUID) (model.Sale, error) {
	sale := model.Sale{}
	if repo.db.Where("id = ?", saleId).Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&sale).Error != nil {
		return sale, fmt.Errorf("Cannot find Sale by id: %w", repo.db.Error)
	}
	return sale, nil
}

func (repo *Repository) GetAllSales() ([]model.Sale, error) {
	sales := make([]model.Sale, 0)
	if repo.db.Preload("Client").Preload("Invoice").Preload("Invoice.Client").Find(&sales).Error != nil {
		return sales, fmt.Errorf("could not get Sales from repo.db: %w", repo.db.Error)
	}
	return sales, nil
}

func (repo *Repository) GetNotInvoicedSalesCountByClientId(clientId uuid.UUID) int64 {
	sales := make([]model.Sale, 0)
	result := repo.db.Where("invoice_id IS NULL AND client_id = ?", clientId).Find(&sales)
	return result.RowsAffected
}

func (repo *Repository) GetSalesByInvoiceId(invoiceId uuid.UUID) ([]model.Sale, error) {
	sales := make([]model.Sale, 0)
	if repo.db.Where("invoice_id = ?", invoiceId).Preload("Client").Find(&sales).Error != nil {
		return sales, fmt.Errorf("could not retrieve sales from repo.db: %w", repo.db.Error)
	}
	return sales, nil
}

func (repo *Repository) UpdateSalesByClientId(clientId uuid.UUID, invoiceId uuid.UUID) (model.Sale, error) {
	sale := model.Sale{InvoiceID: &invoiceId}
	if repo.db.Model(&sale).Where("client_id = ? AND invoice_id IS NULL", clientId).Select("invoice_id").Updates(sale).Error != nil {
		return sale, fmt.Errorf("could not update sales by client id: %w", repo.db.Error)
	}
	return sale, nil
}

func (repo *Repository) CreateSale(sale *model.Sale) (uuid.UUID, error) {
	if repo.db.Create(sale).Error != nil {
		return sale.ID, fmt.Errorf("cannot create Sale: %w", repo.db.Error)
	}
	return sale.ID, nil
}

func (repo *Repository) UpdateSale(sale *model.Sale) error {
	if repo.db.Model(&sale).Select("description", "client_id", "units", "unit_cost", "total").Updates(sale).Error != nil {
		return fmt.Errorf("cannot update Sale: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) DeleteSale(sale *model.Sale) error {
	if repo.db.Delete(sale).Error != nil {
		return fmt.Errorf("cannot delete Sale: %w", repo.db.Error)
	}
	return nil
}
