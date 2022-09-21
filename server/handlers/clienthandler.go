package handlers

import (
	"encoding/json"
	"net/http"
	"server/model"
	"server/request"
	"server/response"
	"server/validation"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (handler *Handler) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := handler.repository.GetAllClients()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving clients", w)
		return
	}
	response.NewOkResponse(&clients, w)
}

func (handler *Handler) GetDeletedClients(w http.ResponseWriter, r *http.Request) {
	clients, err := handler.repository.FindDeletedClients()
	if err != nil {
		response.NewErrorResponse(500, "error occurred retrieving clients", w)
		return
	}
	response.NewOkResponse(&clients, w)
}

func (handler *Handler) GetDeletedClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	client, err := handler.repository.FindDeletedClientByID(clientId)
	if err != nil {
		response.NewErrorResponse(404, "client not found", w)
		return
	}
	response.NewOkResponse(&client, w)
}

func (handler *Handler) GetClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	client, err := handler.repository.FindClientByID(clientId)
	if err != nil {
		response.NewErrorResponse(404, "client not found", w)
		return
	}
	response.NewOkResponse(&client, w)
}

func (handler *Handler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var clientReq request.ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		response.NewErrorResponse(400, "client decode malfunction", w)
		return
	}
	client, err := validation.ClientValidation(&clientReq); if err != nil {
		response.NewErrorResponse(422, "client validation error", w)
		return
	}
	clientId, err := handler.repository.CreateClient(&client); if err != nil {
		response.NewErrorResponse(500, "error occurred creating client", w)
		return
	}
	client.ID = clientId
	response.NewOkResponse(&client, w)
}

func (handler *Handler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	client, err := handler.repository.FindClientByID(clientId)
	if err != nil {
		response.NewErrorResponse(404, "client not found", w)
		return
	}
	clientReq := request.ClientRequest{}
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		response.NewErrorResponse(400, "error occurred decoding client", w)
		return
	}
	clientValidated, err := validation.ClientValidation(&clientReq); if err != nil {
		response.NewErrorResponse(422, "client validation error", w)
		return
	}
	client.Name = clientValidated.Name
	client.Email = clientValidated.Email
	client.Address = clientValidated.Address
	if err := handler.repository.UpdateClient(&client); err != nil {
		response.NewErrorResponse(500, "error occurred updating client", w)
		return
	}
	response.NewOkResponse(&client, w)
}

func (handler *Handler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	_, err = handler.repository.FindClientByID(clientId)
	if err != nil {
		response.NewErrorResponse(404, "client not found", w)
		return
	}
	query := model.Client{}
	query.ID = clientId
	err = handler.repository.DeleteClient(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}

func (handler *Handler) UnDeleteClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	client := model.Client{}
	client.ID = clientId
	if err := handler.repository.UnDeleteClient(&client); err != nil {
		response.NewErrorResponse(500, "error undeleting client", w)
		return 
	}
	response.NewOkResponse(&client, w)
}

func (handler *Handler) PermaDeleteClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		response.NewErrorResponse(400, "invalid uuid", w)
		return
	}
	query := model.Client{}
	query.ID = clientId
	err = handler.repository.PermDeleteClient(&query); if err != nil {
		response.NewErrorResponse(500, "invalid uuid", w)
		return
	}
	response.NewOkResponse(&query, w)
}
