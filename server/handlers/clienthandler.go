package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/model"
	"server/repository"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := repository.GetAllClients()
	if err != nil {
		log.Println("error occurred getting clients")
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&clients)
}

func GetClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", clientIdParam)
		w.WriteHeader(400)
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		log.Printf("client: %v, not found", clientIdParam)
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&client)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	json.NewDecoder(r.Body).Decode(&client)
	clientId, err := repository.CreateClient(&client) 
	if err != nil {
		log.Printf("client: %v, not created", clientId)
		w.WriteHeader(400)
		return
	}
	client.ID = clientId
	json.NewEncoder(w).Encode(&client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", clientIdParam)
		w.WriteHeader(400)
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		log.Printf("client not found with uuid: %v", clientId)
		w.WriteHeader(400)
		return
	}
	req := model.Client{}
	json.NewDecoder(r.Body).Decode(&req)
	client.Name = req.Name
	client.Email = req.Email
	client.Address = req.Address
	
	if err := repository.UpdateClient(&client); err != nil {
		log.Printf("error occurred updating client id: %v", clientId)
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&client)
}
