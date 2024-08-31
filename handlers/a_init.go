package handlers

import (
	r "go-api-template/repositories"
	s "go-api-template/services"
)

type Handler struct {
	repos    *r.Repositories
	services *s.Services
}

func NewHandler(repos *r.Repositories, services *s.Services) *Handler {
	return &Handler{repos: repos, services: services}
}
