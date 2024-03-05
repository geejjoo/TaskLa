package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Storage
	Reserve
	Product
}

func NewRepository(
	db *sqlx.DB,
	storageTableName string,
	reserveTableName string,
	productTableName string,
) *Repository {
	return &Repository{
		NewStorageDB(db, storageTableName),
		NewReserveDB(db, reserveTableName),
		NewProductDB(db, productTableName),
	}
}
