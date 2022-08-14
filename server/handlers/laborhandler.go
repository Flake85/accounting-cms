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
		response.NewErrorResponse(500, "error occurred retrieving labors", w)
		return
	}
	response.NewOkResponse(&labors, w)
}

func GetLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		response.NewErrorResponse(404, "labor not found", w)
		return
	}
	response.NewOkResponse(&labor, w)
}

func CreateLabor(w http.ResponseWriter, r *http.Request) {
	var laborReq request.LaborRequest
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		response.NewErrorResponse(400, "labor decode malfunction", w)
		return
	}
	labor, err := validation.LaborValidation(&laborReq); if err != nil {
		response.NewErrorResponse(422, "labor validation error", w)
		return
	}
	laborId, err := repository.CreateLabor(&labor); if err != nil {
		response.NewErrorResponse(500, "error occurred creating labor", w)
		return
	}
	labor.ID = laborId
	response.NewOkResponse(&labor, w)
}

func UpdateLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	labor, err := repository.FindLaborByID(laborId)
	if err != nil {
		response.NewErrorResponse(404, "labor not found", w)
		return
	}
	laborReq := request.LaborRequest{}
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		response.NewErrorResponse(400, "error occurred decoding labor", w)
		return
	}
	laborValidated, err := validation.LaborValidation(&laborReq); if err != nil {
		response.NewErrorResponse(422, "labor validation error", w)
		return
	}
	labor.Description = laborValidated.Description
	labor.ClientId = laborValidated.ClientId
	labor.InvoiceId = laborValidated.InvoiceId
	labor.HoursWorked = laborValidated.HoursWorked
	labor.HourlyRate = laborValidated.HourlyRate

	if err := repository.UpdateLabor(&labor); err != nil {
		response.NewErrorResponse(500, "error occurred updating labor", w)
		return
	}
	response.NewOkResponse(&labor, w)
}

func DeleteLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	query := model.Labor{}
	query.ID = laborId
	err = repository.DeleteLabor(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
