package repository

import (
	"context"
	"github.com/geejjoo/lamoda_task/cmd/db"
)

type Storage interface {
	StorageInformation(storage string) ([]db.Product, error)
	//ChangeAvailable(storage *db.Storage) error
}

type Reserve interface {
	MakeReservation(ctx context.Context, product *db.Product) error
	CancelReservation(ctx context.Context, product *db.CancelReservation) error
}

type Product interface {
}
