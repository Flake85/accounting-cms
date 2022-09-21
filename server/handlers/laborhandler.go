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

func (handler *Handler) GetLabors(w http.ResponseWriter, r *http.Request) {
	labors, err := handler.repository.GetAllLabors()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving labors", w)
		return
	}
	response.NewOkResponse(&labors, w)
}

func (handler *Handler) GetLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	labor, err := handler.repository.FindLaborByID(laborId)
	if err != nil {
		response.NewErrorResponse(404, "labor not found", w)
		return
	}
	response.NewOkResponse(&labor, w)
}

func (handler *Handler) CreateLabor(w http.ResponseWriter, r *http.Request) {
	var laborReq request.LaborRequest
	if err := json.NewDecoder(r.Body).Decode(&laborReq); err != nil {
		response.NewErrorResponse(400, "labor decode malfunction", w)
		return
	}
	labor, err := validation.LaborValidation(&laborReq); if err != nil {
		response.NewErrorResponse(422, "labor validation error", w)
		return
	}
	client, err := handler.repository.FindClientByID(labor.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	total := labor.HoursWorked * labor.HourlyRate
	labor.Total = math.Round(total * 100) / 100
	
	laborId, err := handler.repository.CreateLabor(&labor); if err != nil {
		response.NewErrorResponse(500, "error occurred creating labor", w)
		return
	}
	labor.ID = laborId
	labor.Client = client
	response.NewOkResponse(&labor, w)
}

func (handler *Handler) UpdateLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	labor, err := handler.repository.FindLaborByID(laborId)
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
	labor.ClientID = laborValidated.ClientID
	labor.HoursWorked = laborValidated.HoursWorked
	labor.HourlyRate = laborValidated.HourlyRate
	total := labor.HoursWorked * labor.HourlyRate
	labor.Total = math.Round(total * 100) / 100

	client, err := handler.repository.FindClientByID(labor.ClientID); if err != nil {
		response.NewErrorResponse(500, "error occured finding client", w)
		return
	}
	if err := handler.repository.UpdateLabor(&labor); err != nil {
		response.NewErrorResponse(500, "error occurred updating labor", w)
		return
	}
	labor.Client = client
	response.NewOkResponse(&labor, w)
}

func (handler *Handler) DeleteLabor(w http.ResponseWriter, r *http.Request) {
	laborIdParam := mux.Vars(r)["id"]
	laborId, err := uuid.Parse(laborIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	labor, err := handler.repository.FindLaborByID(laborId)
	if err != nil {
		response.NewErrorResponse(404, "labor not found", w)
		return
	}
	if labor.InvoiceID != nil {
		response.NewErrorResponse(418, "cannot delete an invoiced labor", w)
		return
	}
	query := model.Labor{}
	query.ID = laborId
	err = handler.repository.DeleteLabor(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
