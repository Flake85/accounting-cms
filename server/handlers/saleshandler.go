package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"server/model"
	"server/request"
	"server/response"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (handler *Handler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := handler.repository.GetAllSales()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieveing sales", w)
		return
	}
	response.NewOkResponse(&sales, w)
}

func (handler *Handler) GetSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	sale, err := handler.repository.FindSaleByID(saleId)
	if err != nil {
		response.NewErrorResponse(404, "sale not found", w)
		return
	}
	response.NewOkResponse(&sale, w)
}

func (handler *Handler) CreateSale(w http.ResponseWriter, r *http.Request) {
	var saleReq request.SaleRequest
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		response.NewErrorResponse(400, "sale decode malfunction", w)
		return
	}
	sale, err := validation.SaleValidation(&saleReq); if err != nil {
		response.NewErrorResponse(422, "sale validation error", w)
		return
	}
	client, err := handler.repository.FindClientByID(sale.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	total := sale.Units * sale.UnitCost
	sale.Total = math.Round(total * 100) / 100

	saleId, err := handler.repository.CreateSale(&sale); if err != nil {
		response.NewErrorResponse(500, "error occurred creating sale", w)
		return
	}
	sale.ID = saleId
	sale.Client = client
	response.NewOkResponse(&sale, w)
}

func (handler *Handler) UpdateSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	sale, err := handler.repository.FindSaleByID(saleId)
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
	sale.Description = saleValidated.Description
	sale.Units = saleValidated.Units
	sale.UnitCost = saleValidated.UnitCost

	client, err := handler.repository.FindClientByID(sale.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}

	total := sale.Units * sale.UnitCost
	sale.Total = math.Round(total * 100) / 100
	
	if err := handler.repository.UpdateSale(&sale); err != nil {
		response.NewErrorResponse(500, "error occurred updating sale", w)
		return
	}
	sale.Client = client
	response.NewOkResponse(&sale, w)
}

func (handler *Handler) DeleteSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	sale, err := handler.repository.FindSaleByID(saleId) 
	if err != nil {
		response.NewErrorResponse(404, "sale not found", w)
		return
	}
	if sale.InvoiceID != nil {
		response.NewErrorResponse(418, "cannot delete an invoiced sale", w)
		return
	}
	query := model.Sale{}
	query.ID = saleId
	err = handler.repository.DeleteSale(&query); if err != nil {
		response.NewErrorResponse(500, "something went wrong", w)
		return
	}
	response.NewOkResponse(&query, w)
}
