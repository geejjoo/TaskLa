package repository

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type ProductDB struct {
	db            *sqlx.DB
	queryReplacer *strings.Replacer
}

func NewProductDB(db *sqlx.DB, tableName string) Product {
	return &ProductDB{
		db:            db,
		queryReplacer: strings.NewReplacer("{table}", tableName),
	}
}
