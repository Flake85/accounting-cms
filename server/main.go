package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db, err := gorm.Open("postgres", "host=localhost port=54321 user=cms dbname=cms password=cms sslmode=disable")
  if err != nil {
	fmt.Printf("db connection error: %v", err)
  }
  defer db.Close()
}