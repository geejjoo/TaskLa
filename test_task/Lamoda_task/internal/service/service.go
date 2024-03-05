package service

import (
	"github.com/geejjoo/lamoda_task/internal/repository"
)

type Service struct {
	Storage
	Reserve
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewStorageService(*repos),
		NewReserveService(*repos),
	}
}
