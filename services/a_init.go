package services

import r "go-api-template/repositories"

type Services struct {
	repos *r.Repositories
}

func InitServices(repos *r.Repositories) *Services {
	return &Services{repos: repos}
}
