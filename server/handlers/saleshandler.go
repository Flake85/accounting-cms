package handlers

import (
	"encoding/json"
	"net/http"
	"server/model"
	"server/repository"
	"server/request"
	"server/response"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := repository.GetAllSales()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieveing sales", w)
		return
	}
	response.NewOkResponse(&sales, w)
}

func GetSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		response.NewErrorResponse(404, "sale not found", w)
		return
	}
	response.NewOkResponse(&sale, w)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	var saleReq request.SaleRequest
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		response.NewErrorResponse(400, "sale decode malfunction", w)
		return
	}
	sale, err := validation.SaleValidation(&saleReq); if err != nil {
		response.NewErrorResponse(422, "sale validation error", w)
		return
	}
	sale.Client, err = repository.FindClientByID(sale.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	if sale.InvoiceID != nil {
		invoice, err := repository.FindInvoiceByID(*sale.InvoiceID); 
		if err != nil {
			response.NewErrorResponse(500, "error occured finding invoice", w)
			return
		}
		if sale.ClientID != invoice.ClientID {
			response.NewErrorResponse(500, "sale's client id must match invoice's client id", w)
			return
		}
		invoice.Client = sale.Client
		sale.Invoice = &invoice
	}
	total := sale.Units * sale.UnitCost
	sale.Total = total
	saleId, err := repository.CreateSale(&sale); if err != nil {
		response.NewErrorResponse(500, "error occurred creating sale", w)
		return
	}
	sale.ID = saleId
	response.NewOkResponse(&sale, w)
}

func UpdateSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		response.NewErrorResponse(404, "sale not found", w)
		return
	}
	saleReq := request.SaleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		response.NewErrorResponse(400, "error occurred decoding sale", w)
		return
	}
	saleValidated, err := validation.SaleValidation(&saleReq); if err != nil {
		response.NewErrorResponse(422, "sale validation error", w)
		return
	}
	sale.ClientID = saleValidated.ClientID
	sale.InvoiceID = saleValidated.InvoiceID
	sale.Description = saleValidated.Description
	sale.Units = saleValidated.Units
	sale.UnitCost = saleValidated.UnitCost
	
	total := sale.Units * sale.UnitCost
	sale.Total = total

	if sale.InvoiceID != nil {
		invoice, err := repository.FindInvoiceByID(*sale.InvoiceID); 
		if err != nil {
			response.NewErrorResponse(500, "error occured finding invoice", w)
			return
		}
		if sale.ClientID != invoice.ClientID {
			response.NewErrorResponse(500, "sale's client id must match invoice's client id", w)
			return
		}
		invoice.Client = sale.Client
		sale.Invoice = &invoice
	}
	sale.Client, err = repository.FindClientByID(sale.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	if err := repository.UpdateSale(&sale); err != nil {
		response.NewErrorResponse(500, "error occurred updating sale", w)
		return
	}
	response.NewOkResponse(&sale, w)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	query := model.Sale{}
	query.ID = saleId
	err = repository.DeleteSale(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
