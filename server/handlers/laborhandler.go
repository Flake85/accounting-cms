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

func GetLabors(w http.ResponseWriter, r *http.Request) {
	labors, err := repository.GetAllLabors()
	if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred retrieving labors"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&labors)
	json.NewEncoder(w).Encode(res.Body)
}

func GetLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("labor not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&labor)
	json.NewEncoder(w).Encode(res.Body)
}

func CreateLabor(w http.ResponseWriter, r *http.Request) {
	var laborReq request.LaborRequest
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("labor decode malfunction"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	labor, err := validation.LaborValidation(&laborReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("labor validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	laborId, err := repository.CreateLabor(&labor); if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred creating labor"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	labor.ID = laborId
	res := response.NewOkResponse(&labor)
	json.NewEncoder(w).Encode(res.Body)
}

func UpdateLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("labor not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	laborReq := request.LaborRequest{}
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("error occurred decoding labor"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	laborValidated, err := validation.LaborValidation(&laborReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("labor validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	labor.Description = laborValidated.Description
	labor.ClientId = laborValidated.ClientId
	labor.InvoiceId = laborValidated.InvoiceId
	labor.HoursWorked = laborValidated.HoursWorked
	labor.HourlyRate = laborValidated.HourlyRate

	if err := repository.UpdateLabor(&labor); err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred updating labor"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&labor)
	json.NewEncoder(w).Encode(res.Body)
}

func DeleteLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	query := model.Labor{}
	query.ID = laborId
	err = repository.DeleteLabor(&query); if err != nil {
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
