package repository

import (
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/jmoiron/sqlx"
	"strings"
)

type StorageDB struct {
	db            *sqlx.DB
	queryReplacer *strings.Replacer
}

func NewStorageDB(db *sqlx.DB, tableName string) Storage {
	return &StorageDB{
		db:            db,
		queryReplacer: strings.NewReplacer("{table}", tableName),
	}
}

func (w *StorageDB) StorageInformation(storage string) ([]db.Product, error) {
	var storageList []db.Product
	query := w.queryReplacer.Replace("SELECT articular, quantity FROM RESERVE WHERE storage=($1)")
	err := w.db.Select(&storageList, query, storage)
	if err != nil {
		return nil, InvalidStorage
	}
	return storageList, nil
}
