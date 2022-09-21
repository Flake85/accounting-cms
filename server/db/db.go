package db

import (
	"fmt"
	"server/config"
	"server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(configuration config.Configuration) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.DbName, configuration.Password,
	)))
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Client{}, &model.Expense{}, &model.Invoice{}, &model.Labor{}, &model.Sale{})
}
