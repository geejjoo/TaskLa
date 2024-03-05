package service

import (
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/geejjoo/lamoda_task/internal/repository"
)

type StorageService struct {
	repo repository.Repository
}

func NewStorageService(repo repository.Repository) Storage {
	return &StorageService{repo: repo}
}

func (w *StorageService) StorageInformation(storage string) ([]db.Product, error) {
	inventory, err := w.repo.StorageInformation(storage)
	if err != nil {
		return nil, StorageInformation
	}

	return inventory, err

}

func (w *StorageService) ChangeAvailable() error {
	return nil
}
