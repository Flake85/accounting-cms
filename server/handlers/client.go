package handlers

import (
	"encoding/json"
	"net/http"
	"server/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetClients(w http.ResponseWriter, r *http.Request) {
	var clients []model.Client
	DB.Find(&clients)
	json.NewEncoder(w).Encode(&clients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	json.NewDecoder(r.Body).Decode(&client)
	DB.Create(&client)
	json.NewEncoder(w).Encode(&client)
}
