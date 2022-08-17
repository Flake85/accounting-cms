package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"server/model"
	"server/repository"
	router "server/router"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=54321 user=cms dbname=cms password=cms sslmode=disable")
  	if err != nil {
		fmt.Printf("db connection error: %v", err)
  	}
  	defer db.Close()
	db.AutoMigrate(&model.Client{})
	repository.DB = db
	r := router.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}