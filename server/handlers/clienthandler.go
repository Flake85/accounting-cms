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

func GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := repository.GetAllClients()
	if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred retrieving clients"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&clients)
	json.NewEncoder(w).Encode(res.Body)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("client not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&client)
	json.NewEncoder(w).Encode(res.Body)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var clientReq request.ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("client decode malfunction"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	client, err := validation.ClientValidation(&clientReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("client validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	clientId, err := repository.CreateClient(&client); if err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred creating client"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	client.ID = clientId
	res := response.NewOkResponse(&client)
	json.NewEncoder(w).Encode(res.Body)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		res := response.NewErrorResponse(
			404, response.NewBaseMessage("client not found"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	clientReq := request.ClientRequest{}
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("error occurred decoding client"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	clientValidated, err := validation.ClientValidation(&clientReq); if err != nil {
		res := response.NewErrorResponse(
			422, response.NewBaseMessage("client validation error"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	client.Name = clientValidated.Name
	client.Email = clientValidated.Email
	client.Address = clientValidated.Address
	if err := repository.UpdateClient(&client); err != nil {
		res := response.NewErrorResponse(
			500, response.NewBaseMessage("error occurred updating client"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	res := response.NewOkResponse(&client)
	json.NewEncoder(w).Encode(res.Body)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		res := response.NewErrorResponse(
			400, response.NewBaseMessage("invalid uuid"),
		)
		w.WriteHeader(res.Code)
		json.NewEncoder(w).Encode(res.Body)
		return
	}
	query := model.Client{}
	query.ID = clientId
	err = repository.DeleteClient(&query); if err != nil {
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
