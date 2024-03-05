package service

import (
	"context"
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/geejjoo/lamoda_task/internal/repository"
	"time"
)

type ReserveService struct {
	repo repository.Repository
}

func NewReserveService(repo repository.Repository) Reserve {
	return &ReserveService{repo: repo}
}

func (w *ReserveService) MakeReservation(product *db.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := w.repo.MakeReservation(ctx, product)

	return err

}

func (w *ReserveService) CancelReservation(product *db.CancelReservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := w.repo.CancelReservation(ctx, product)

	return err
}
