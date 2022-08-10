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

func GetLabors(w http.ResponseWriter, r *http.Request) {
	labors, err := repository.GetAllLabors()
	if err != nil {
		log.Println("error occurred getting labors")
		w.WriteHeader(500)
		fmt.Fprintln(w, "error occurred retrieving labors")
		return
	}
	json.NewEncoder(w).Encode(&labors)
}

func GetLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", laborIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		log.Printf("labor: %v, not found", laborIdParam)
		w.WriteHeader(404)
		fmt.Fprintln(w, "labor not found")
		return
	}
	json.NewEncoder(w).Encode(&labor)
}

func CreateLabor(w http.ResponseWriter, r *http.Request) {
	var laborReq request.LaborRequest
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		log.Printf("labor decode malfunction: %v", err)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred creating labor")
		return
	}
	labor, err := validation.LaborValidation(&laborReq); if err != nil {
		log.Println("validation error. labor not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "labor validation error: %v", err)
		return
	}
	laborId, err := repository.CreateLabor(&labor); if err != nil {
		log.Printf("labor: %v, not created", laborId)
		w.WriteHeader(500)
		fmt.Fprintf(w, "an error occurred creating labor: %v", err)
		return
	}
	labor.ID = laborId
	json.NewEncoder(w).Encode(&labor)
}

func UpdateLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", laborIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating labor")
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		log.Printf("labor not found with uuid: %v", laborId)
		w.WriteHeader(404)
		fmt.Fprintln(w, "labor not found")
		return
	}
	laborReq := request.LaborRequest{}
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		log.Print("labor decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating labor")
		return
	}
	laborValidated, err := validation.LaborValidation(&laborReq); if err != nil {
		log.Println("validation error. labor not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "labor validation error: %v", err)
		return
	}
	labor.Description = laborValidated.Description
	labor.ClientId = laborValidated.ClientId
	labor.InvoiceId = laborValidated.InvoiceId
	labor.HoursWorked = laborValidated.HoursWorked
	labor.HourlyRate = laborValidated.HourlyRate

	if err := repository.UpdateLabor(&labor); err != nil {
		log.Printf("error occurred updating labor id: %v", laborId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred updating labor")
		return
	}
	json.NewEncoder(w).Encode(&labor)
}

func DeleteLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", laborIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	query := model.Labor{}
	query.ID = laborId
	err = repository.DeleteLabor(&query); if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "error deleting labor")
		return
	}
	json.NewEncoder(w).Encode(&query)
}
