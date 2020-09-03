package service

import (
	"elipzis.com/inertia-echo/repository"
)

//
type Service struct {
	repository *repository.Repository
}

// Services know about their repositories
func NewService(repo *repository.Repository) (this *Service) {
	this = new(Service)
	if repo == nil {
		this.repository = repository.NewRepository(repository.DB.Conn)
	} else {
		this.repository = repo
	}
	return this
}
