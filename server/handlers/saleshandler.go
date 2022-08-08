package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/model"
	"server/repository"
	"server/request"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := repository.GetAllSales()
	if err != nil {
		log.Println("error occurred getting sales")
		w.WriteHeader(500)
		fmt.Fprintln(w, "error occurred retrieving sales")
		return
	}
	json.NewEncoder(w).Encode(&sales)
}

func GetSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", saleIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		log.Printf("sale: %v, not found", saleIdParam)
		w.WriteHeader(404)
		fmt.Fprintln(w, "sale not found")
		return
	}
	json.NewEncoder(w).Encode(&sale)
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	var saleReq request.SaleRequest
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		log.Printf("sale decode malfunction: %v", err)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred creating sale")
		return
	}
	sale, err := validation.SaleValidation(&saleReq); if err != nil {
		log.Println("sale not created")
		w.WriteHeader(400)
		fmt.Fprintf(w, "an error occurred creating sale: %v", err)
		return
	}
	saleId, err := repository.CreateSale(&sale); if err != nil {
		log.Printf("sale: %v, not created", saleId)
		w.WriteHeader(500)
		fmt.Fprintf(w, "an error occurred creating sale: %v", err)
		return
	}
	sale.ID = saleId
	json.NewEncoder(w).Encode(&sale)
}

func UpdateSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", saleIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating sale")
		return
	}
	sale, err := repository.FindSaleByID(saleId)
	if err != nil {
		log.Printf("sale not found with uuid: %v", saleId)
		w.WriteHeader(404)
		fmt.Fprintln(w, "sale not found")
		return
	}
	saleReq := request.SaleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&saleReq); err != nil {
		log.Print("sale decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating sale")
		return
	}
	saleValidated, err := validation.SaleValidation(&saleReq); if err != nil {
		log.Println("sale not created")
		w.WriteHeader(400)
		fmt.Fprintf(w, "an error occurred creating sale: %v", err)
		return
	}
	sale.ClientId = saleValidated.ClientId
	sale.InvoiceId = saleValidated.InvoiceId
	sale.Description = saleValidated.Description
	sale.Units = saleValidated.Units
	sale.UnitCost = saleValidated.UnitCost

	if err := repository.UpdateSale(&sale); err != nil {
		log.Printf("error occurred updating sale id: %v", saleId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred updating sale")
		return
	}
	json.NewEncoder(w).Encode(&sale)
}

func DeleteSale(w http.ResponseWriter, r *http.Request) {
	saleIdParam := mux.Vars(r)["id"]
	saleId, err := uuid.Parse(saleIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", saleIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	query := model.Sale{}
	query.ID = saleId
	err = repository.DeleteSale(&query); if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	json.NewEncoder(w).Encode(&query)
}