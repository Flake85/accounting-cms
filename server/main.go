package main

import (
	"fmt"
	"log"
	"net/http"

	"server/config"
	"server/db"
	"server/handlers"
	"server/repository"
	"server/router"
)

func main() {
	cfg := config.Parse()
	database, err := db.NewDB(cfg); if err != nil {
		fmt.Printf("db connection error: %v", err)
	}
	db.Migrate(database)
	repo := repository.NewRepository(database)
	handler := handlers.NewHandler(&repo)
	
	r := router.NewRouter(&handler)
	log.Println("server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
