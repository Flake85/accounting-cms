package handlers

import (
	"encoding/json"
	"net/http"
	"server/repository"
	"server/response"

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
		response.NewErrorResponse(404, "client not found", w)
		return
	}
	response.NewOkResponse(&invoice, w)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		response.NewErrorResponse(400, "invoice decode malfunction", w)
		return
	}
	invoiceId, err := repository.CreateInvoice(&invoice) 
	if err != nil {
		response.NewErrorResponse(500, "error occurred creating invoice", w)
		return
	}
	invoice.ID = invoiceId
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
	invoice.IsInvoiced = req.IsInvoiced
	
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
	query := model.Invoice{}
	query.ID = invoiceId
	err = repository.DeleteInvoice(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
