package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/repository"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"server/model"
)

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	invoices, err := repository.GetAllInvoices()
	if err != nil {
		log.Println("error occurred getting invoices")
		w.WriteHeader(500)
		fmt.Fprintln(w, "error occurred retrieving invoices")
		return
	}
	json.NewEncoder(w).Encode(&invoices)
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", invoiceIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
	if err != nil {
		log.Printf("invoice: %v, not found", invoiceIdParam)
		w.WriteHeader(404)
		fmt.Fprintln(w, "invoice not found")
		return
	}
	json.NewEncoder(w).Encode(&invoice)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice model.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		log.Print("invoice decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred creating invoice")
		return
	}
	invoiceId, err := repository.CreateInvoice(&invoice) 
	if err != nil {
		log.Printf("invoice: %v, not created", invoiceId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred creating invoice")
		return
	}
	invoice.ID = invoiceId
	json.NewEncoder(w).Encode(&invoice)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", invoiceIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating invoice")
		return
	}
	invoice, err := repository.FindInvoiceByID(invoiceId)
	if err != nil {
		log.Printf("invoice not found with uuid: %v", invoiceId)
		w.WriteHeader(404)
		fmt.Fprintln(w, "invoice not found")
		return
	}
	req := model.Invoice{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Print("invoice decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating invoice")
		return
	}
	invoice.IsInvoiced = req.IsInvoiced
	
	if err := repository.UpdateInvoice(&invoice); err != nil {
		log.Printf("error occurred updating invoice id: %v", invoiceId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred updating invoice")
		return
	}
	json.NewEncoder(w).Encode(&invoice)
}

func DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	invoiceIdParam := mux.Vars(r)["id"]
	invoiceId, err := uuid.Parse(invoiceIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", invoiceIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	query := model.Invoice{}
	query.ID = invoiceId
	err = repository.DeleteInvoice(&query); if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	json.NewEncoder(w).Encode(&query)
}
