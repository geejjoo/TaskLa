package repository

import (
	"context"
	"database/sql"
	"github.com/geejjoo/lamoda_task/cmd/db"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ReserveDB struct {
	db            *sqlx.DB
	queryReplacer *strings.Replacer
}

func NewReserveDB(db *sqlx.DB, tableName string) Reserve {
	return &ReserveDB{
		db:            db,
		queryReplacer: strings.NewReplacer("{table}", tableName),
	}
}

func (w *ReserveDB) MakeReservation(ctx context.Context, product *db.Product) error {
	tx, err := w.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}

	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT * FROM {table} WHERE articular = ($1) LIMIT 1", product.Articular).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return NoArticular
	}

	query := w.queryReplacer.Replace("SELECT * FROM {table} WHERE articular=($1) and quantity>=($2)")
	row := tx.QueryRow(query, product.Articular, product.Quantity)
	if err = row.Scan(&product.Articular, &product.Quantity); err != nil {
		return NotEnoughQuantity
	}

	query = w.queryReplacer.Replace("UPDATE {table} SET quantity = quantity - ($1), reserved = ($2) WHERE articular=($3) and quantity>=($4) and storage = (SELECT storage FROM {table} WHERE articular=($5) and quantity>=($6) LIMIT 1)")
	_, err = tx.Exec(query, product.Quantity, product.Quantity, product.Articular, product.Quantity, product.Articular, product.Quantity)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (w *ReserveDB) CancelReservation(ctx context.Context, product *db.CancelReservation) error {
	tx, err := w.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	if err != nil {
		return err
	}

	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM {table} WHERE articular = $1)", product.Articular).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return NoArticular
	}

	query := w.queryReplacer.Replace("SELECT * FROM {table} WHERE articular=($1) and reserved<=($2) and storage=($3)")
	row := tx.QueryRow(query, product.Articular, product.Quantity, product.Storage)
	if err = row.Scan(&product.Articular, &product.Quantity, &product.Storage); err != nil {
		return NoReservation
	}

	query = w.queryReplacer.Replace("UPDATE {table} SET quantity = quantity + ($1), reserved = reserved - ($2) WHERE articular=($3) and reserved>=($4) and storage = ($5)")
	_, err = tx.Exec(query, product.Quantity, product.Quantity, product.Articular, product.Quantity, product.Storage)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
