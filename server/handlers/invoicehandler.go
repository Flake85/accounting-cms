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
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred retrieving invoices"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&invoices)
	json.NewEncoder(w).Encode(res.Body)
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("client not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&invoice)
	json.NewEncoder(w).Encode(res.Body)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invoice decode malfunction"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	invoiceId, err := repository.CreateInvoice(&invoice) 
	if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred creating invoice"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	invoice.ID = invoiceId
	res := response.NewOkResponse(&invoice)
	json.NewEncoder(w).Encode(res.Body)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("error occurred creating invoice"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("invoice not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	req := model.Invoice{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invoice decode malfunction"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	invoice.IsInvoiced = req.IsInvoiced
	
	if err := repository.UpdateInvoice(&invoice); err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred updating expense"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&invoice)
	json.NewEncoder(w).Encode(res.Body)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	query := model.Invoice{}
	query.ID = invoiceId
	err = repository.DeleteInvoice(&query); if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&query)
	json.NewEncoder(w).Encode(res.Body)
}
