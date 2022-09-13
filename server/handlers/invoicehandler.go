package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"server/repository"
	"server/response"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"server/model"
)

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := repository.GetAllInvoices()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving invoices", w)
		return
	}
	response.NewOkResponse(&invoices, w)
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "invoice not found", w)
		return
	}
	sales, err := repository.GetSalesByInvoiceId(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "sales not found", w)
		return
	}
	for _, sale := range sales {
		invoice.SalesTotal += sale.Total
	}
	labors, err := repository.GetLaborsByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(404, "labors not found", w)
		return
	}
	for _, labor := range labors {
		invoice.LaborsTotal += labor.Total
	}
	invoice.LaborsTotal = math.Round(invoice.LaborsTotal * 100) / 100
	invoice.GrandTotal = invoice.SalesTotal + invoice.LaborsTotal
	invoice.Sales = &sales
	invoice.Labors = &labors
	response.NewOkResponse(&invoice, w)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		response.NewErrorResponse(400, "invoice decode malfunction", w)
		return
	}
	client, err := repository.FindClientByID(invoice.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}

	laborCount := repository.GetNotInvoicedLaborsCountByClientId(invoice.ClientID)
	salesCount := repository.GetNotInvoicedSalesCountByClientId(invoice.ClientID)
	if laborCount < 1 && salesCount < 1{
		response.NewErrorResponse(400, "cannot create empty invoice. add sale or labor first.", w)
		return
	}

	current := time.Now().Format("2006-01-02")
	description := client.Name + ": " + current 
	invoice.Description = description

	invoiceId, err := repository.CreateInvoice(&invoice) 
	if err != nil {
		response.NewErrorResponse(500, "error occurred creating invoice", w)
		return
	}
	invoice.ID = invoiceId
	invoice.Client = client
	_, err = repository.UpdateSalesByClientId(invoice.ClientID, invoiceId); if err != nil {
		response.NewErrorResponse(500, "error updating invoice sales with client id", w)
		return
	}
	_, err = repository.UpdateLaborsByClientId(invoice.ClientID, invoiceId); if err != nil {
		response.NewErrorResponse(500, "error updating invoice labors with client id", w)
		return
	}
	sales, err := repository.GetSalesByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(500, "error retrieveing sales by invoice id", w)
		return
	}
	invoice.Sales = &sales
	labors, err := repository.GetLaborsByInvoiceId(invoiceId); if err != nil {
		response.NewErrorResponse(500, "error retrieving labors by invoice id", w)
		return
	}
	invoice.Labors = &labors
	response.NewOkResponse(&invoice, w)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "error occurred creating invoice", w)
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
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
	invoice.Client, err = repository.FindClientByID(invoice.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	if err := repository.UpdateInvoice(&invoice); err != nil {
		response.NewErrorResponse(500, "error occurred updating expense", w)
		return
	}
	response.NewOkResponse(&invoice, w)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	_, err = repository.FindInvoiceByID(invoiceId)
	if err != nil {
		response.NewErrorResponse(404, "invoice not found", w)
		return
	}
	query := model.Invoice{}
	query.ID = invoiceId
	err = repository.DeleteInvoice(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
