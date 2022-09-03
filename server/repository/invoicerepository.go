package repository

import (
	"fmt"
	"server/model"

	"github.com/google/uuid"
)

func FindInvoiceByID(invoiceId uuid.UUID) (model.Invoice, error) {
	invoice := model.Invoice{}
	if DB.Where("id = ?", invoiceId).Preload("Client").Find(&invoice).Error != nil {
		return invoice, fmt.Errorf("cannot find invoice by id: %w", DB.Error)
	}
	return invoice, nil
}

func GetAllInvoices() ([]model.Invoice, error) {
	invoices := make([]model.Invoice, 0)
	if DB.Preload("Client").Find(&invoices).Error != nil {
		return invoices, fmt.Errorf("could not get invoices from db: %w", DB.Error)
	}
	return invoices, nil
}

func CreateInvoice(invoice *model.Invoice) (uuid.UUID, error) {
	if DB.Create(invoice).Error != nil {
		return invoice.ID, fmt.Errorf("cannot create invoice: %w", DB.Error)
	}
	return invoice.ID, nil
}

func UpdateInvoice(invoice *model.Invoice) error {
	if DB.Save(invoice).Error != nil {
		return fmt.Errorf("cannot update invoice: %w", DB.Error)
	}
	return nil
}

func DeleteInvoice(invoice *model.Invoice) error {
	if DB.Delete(invoice).Error != nil {
		return fmt.Errorf("cannot delete invoice: %w", DB.Error)
	}
	return nil
}