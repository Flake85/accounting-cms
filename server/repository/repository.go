package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return Repository{ db: database }
}
