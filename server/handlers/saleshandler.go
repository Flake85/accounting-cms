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
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred retrieveing sales"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&sales)
	json.NewEncoder(w).Encode(res.Body)
}

func GetSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("sale not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&sale)
	json.NewEncoder(w).Encode(res.Body)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	var saleReq request.SaleRequest
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("sale decode malfunction"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	sale, err := validation.SaleValidation(&saleReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("sale validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	saleId, err := repository.CreateSale(&sale); if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred creating sale"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	sale.ID = saleId
	res := response.NewOkResponse(&sale)
	json.NewEncoder(w).Encode(res.Body)
}

func UpdateSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("sale not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	saleReq := request.SaleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("error occurred decoding sale"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	saleValidated, err := validation.SaleValidation(&saleReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("sale validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	sale.ClientId = saleValidated.ClientId
	sale.InvoiceId = saleValidated.InvoiceId
	sale.Description = saleValidated.Description
	sale.Units = saleValidated.Units
	sale.UnitCost = saleValidated.UnitCost

	if err := repository.UpdateSale(&sale); err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred updating sale"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&sale)
	json.NewEncoder(w).Encode(res.Body)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	query := model.Sale{}
	query.ID = saleId
	err = repository.DeleteSale(&query); if err != nil {
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
