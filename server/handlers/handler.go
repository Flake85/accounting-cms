package handlers

import "server/repository"

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repo *repository.Repository) Handler {
	return Handler{ repository: repo }
}
