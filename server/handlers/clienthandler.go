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

func GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := repository.GetAllClients()
	if err != nil {
		log.Println("error occurred getting clients")
		w.WriteHeader(500)
		fmt.Fprintln(w, "error occurred retrieving clients")
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
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		log.Printf("client: %v, not found", clientIdParam)
		w.WriteHeader(404)
		fmt.Fprintln(w, "client not found")
		return
	}
	json.NewEncoder(w).Encode(&client)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var clientReq request.ClientRequest
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		log.Print("client decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred creating client")
		return
	}
	client, err := validation.ClientValidation(&clientReq); if err != nil {
		log.Println("validation error. client not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "client validation error: %v", err)
		return
	}
	clientId, err := repository.CreateClient(&client); if err != nil {
		log.Printf("client: %v, not created", clientId)
		w.WriteHeader(500)
		fmt.Fprintf(w, "an error occurred creating client: %v", err)
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
		fmt.Fprintln(w, "an error occurred updating client")
		return
	}
	client, err := repository.FindClientByID(clientId)
	if err != nil {
		log.Printf("client not found with uuid: %v", clientId)
		w.WriteHeader(404)
		fmt.Fprintln(w, "client not found")
		return
	}
	clientReq := request.ClientRequest{}
	if err := json.NewDecoder(r.Body).Decode(&clientReq); err != nil {
		log.Print("client decode malfunction")
		w.WriteHeader(400)
		fmt.Fprintln(w, "an error occurred updating client")
		return
	}
	clientValidated, err := validation.ClientValidation(&clientReq); if err != nil {
		log.Println("validation error. client not created")
		w.WriteHeader(422)
		fmt.Fprintf(w, "client validation error: %v", err)
		return
	}
	client.Name = clientValidated.Name
	client.Email = clientValidated.Email
	client.Address = clientValidated.Address
	if err := repository.UpdateClient(&client); err != nil {
		log.Printf("error occurred updating client id: %v", clientId)
		w.WriteHeader(500)
		fmt.Fprintln(w, "an error occurred updating client")
		return
	}
	json.NewEncoder(w).Encode(&client)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	clientIdParam := mux.Vars(r)["id"]
	clientId, err := uuid.Parse(clientIdParam)
	if err != nil {
		log.Printf("%v, is not a valid uuid.", clientIdParam)
		w.WriteHeader(400)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	query := model.Client{}
	query.ID = clientId
	err = repository.DeleteClient(&query); if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "invalid uuid")
		return
	}
	json.NewEncoder(w).Encode(&query)
}
