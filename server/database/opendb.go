package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/autumnleaf-ra/go-movie-api/models"
	_ "github.com/lib/pq"
)

func OpenDB(cfg models.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DB.DSN)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
