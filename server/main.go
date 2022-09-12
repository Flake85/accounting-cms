package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"server/db"
	"server/flags"
	"server/model"
	"server/repository"
	"server/router"
)

func main() {
	flag.Parse()
	db, err := gorm.Open(postgres.Open(db.BuildDbConnectionStr(*flags.Host, *flags.Port, *flags.User, *flags.Password, *flags.DbName)))
	if err != nil {
		fmt.Printf("db connection error: %v", err)
	}
	repository.DB = db
	db.AutoMigrate(&model.Client{}, &model.Expense{}, &model.Invoice{}, &model.Labor{}, &model.Sale{})
	r := router.NewRouter()
	log.Println("server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
