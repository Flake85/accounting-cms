package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func (repo *Repository) FindInvoiceByID(invoiceId uuid.UUID) (model.Invoice, error) {
	invoice := model.Invoice{}
	if repo.db.Where("id = ?", invoiceId).Preload("Client").Find(&invoice).Error != nil {
		return invoice, fmt.Errorf("cannot find invoice by id: %w", repo.db.Error)
	}
	return invoice, nil
}

func (repo *Repository) GetAllInvoices() ([]model.Invoice, error) {
	invoices := make([]model.Invoice, 0)
	if repo.db.Preload("Client").Find(&invoices).Error != nil {
		return invoices, fmt.Errorf("could not get invoices from repo.db: %w", repo.db.Error)
	}
	return invoices, nil
}

func (repo *Repository) CreateInvoice(invoice *model.Invoice) (uuid.UUID, error) {
	if repo.db.Create(invoice).Error != nil {
		return invoice.ID, fmt.Errorf("cannot create invoice: %w", repo.db.Error)
	}
	return invoice.ID, nil
}

func (repo *Repository) UpdateInvoice(invoice *model.Invoice) error {
	if repo.db.Model(&invoice).Select("is_paid").Updates(invoice).Error != nil {
		return fmt.Errorf("cannot update invoice: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) UpdateInvoiceTotals(invoice *model.Invoice) error {
	if repo.db.Model(&invoice).Select("sales_total", "labors_total", "grand_total").Updates(invoice).Error != nil {
		return fmt.Errorf("cannot update totals: %w", repo.db.Error)
	}
	return nil
}

func (repo *Repository) DeleteInvoice(invoice *model.Invoice) error {
	if repo.db.Delete(invoice).Error != nil {
		return fmt.Errorf("cannot delete invoice: %w", repo.db.Error)
	}
	return nil
}
