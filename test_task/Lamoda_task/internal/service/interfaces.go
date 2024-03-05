package service

import "github.com/geejjoo/lamoda_task/cmd/db"

type Storage interface {
	StorageInformation(storage string) ([]db.Product, error)
	ChangeAvailable() error
}

type Reserve interface {
	MakeReservation(product *db.Product) error
	CancelReservation(product *db.CancelReservation) error
}
