package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DriverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, DatabaseError
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
