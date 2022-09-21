package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"server/response"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"server/model"
)

func (handler *Handler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := handler.repository.GetAllInvoices()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving invoices", w)
		return
	}
	response.NewOkResponse(&invoices, w)
}

func (handler *Handler) GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	invoice, err := handler.repository.FindInvoiceByID(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "invoice not found", w)
		return
	}
	sales, err := handler.repository.GetSalesByInvoiceId(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "sales not found", w)
		return
	}
	labors, err := handler.repository.GetLaborsByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(404, "labors not found", w)
		return
	}
	invoice.Sales = &sales
	invoice.Labors = &labors
	response.NewOkResponse(&invoice, w)
}

func (handler *Handler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		response.NewErrorResponse(400, "invoice decode malfunction", w)
		return
	}
	client, err := handler.repository.FindClientByID(invoice.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}

	laborCount := handler.repository.GetNotInvoicedLaborsCountByClientId(invoice.ClientID)
	salesCount := handler.repository.GetNotInvoicedSalesCountByClientId(invoice.ClientID)
	if laborCount < 1 && salesCount < 1{
		response.NewErrorResponse(400, "cannot create empty invoice. add sale or labor first.", w)
		return
	}

	current := time.Now().Format("2006-01-02")
	description := client.Name + ": " + current 
	invoice.Description = description

	invoiceId, err := handler.repository.CreateInvoice(&invoice) 
	if err != nil {
		response.NewErrorResponse(500, "error occurred creating invoice", w)
		return
	}
	invoice.ID = invoiceId
	invoice.Client = client
	_, err = handler.repository.UpdateSalesByClientId(invoice.ClientID, invoiceId); if err != nil {
		response.NewErrorResponse(500, "error updating invoice sales with client id", w)
		return
	}
	_, err = handler.repository.UpdateLaborsByClientId(invoice.ClientID, invoiceId); if err != nil {
		response.NewErrorResponse(500, "error updating invoice labors with client id", w)
		return
	}
	sales, err := handler.repository.GetSalesByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(500, "error retrieveing sales by invoice id", w)
		return
	}
	labors, err := handler.repository.GetLaborsByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(500, "error retrieving labors by invoice id", w)
		return
	}
	invoice.Sales = &sales
	invoice.Labors = &labors
	for _, sale := range sales {
		invoice.SalesTotal += sale.Total
	}
	for _, labor := range labors {
		invoice.LaborsTotal += labor.Total
	}
	laborsTotal := math.Round(invoice.LaborsTotal * 100) / 100
	grandTotal := invoice.SalesTotal + invoice.LaborsTotal
	invoice.LaborsTotal = laborsTotal
	invoice.GrandTotal = grandTotal
	err = handler.repository.UpdateInvoiceTotals(&invoice); if err != nil {
		response.NewErrorResponse(500, "error saving invoice totals", w)
		return
	}
	response.NewOkResponse(&invoice, w)
}

func (handler *Handler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "error occurred creating invoice", w)
		return
	}
	invoice, err := handler.repository.FindInvoiceByID(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "invoice not found", w)
		return
	}
	req := model.Invoice{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewErrorResponse(400, "invoice decode malfunction", w)
		return
	}
	invoice.IsPaid = req.IsPaid
	client, err := handler.repository.FindClientByID(invoice.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	if err := handler.repository.UpdateInvoice(&invoice); err != nil {
		response.NewErrorResponse(500, "error occurred updating expense", w)
		return
	}
	invoice.Client = client
	response.NewOkResponse(&invoice, w)
}

func (handler *Handler) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	_, err = handler.repository.FindInvoiceByID(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "invoice not found", w)
		return
	}
	query := model.Invoice{}
	query.ID = invoiceId
	err = handler.repository.DeleteInvoice(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
