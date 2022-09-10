package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"server/db"
	"server/flags"
	"server/model"
	"server/repository"
	"server/router"
)

func main() {
	flag.Parse()
	db, err := gorm.Open("postgres", db.BuildDbConnectionStr(*flags.Host, *flags.Port, *flags.User, *flags.Password, *flags.DbName))
	if err != nil {
		fmt.Printf("db connection error: %v", err)
	}
	defer db.Close()
	repository.DB = db
	db.AutoMigrate(&model.Client{}, &model.Expense{}, &model.Invoice{}, &model.Labor{}, &model.Sale{})
	r := router.NewRouter()
	log.Println("server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
