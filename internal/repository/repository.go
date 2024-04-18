package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Eiphoria/GoReversi/internal/config"
)

type Repository struct {
	db *sql.DB
}

func New(cfg config.DBConfig) (*Repository, error) {
	dbs, err := sql.Open("postgres", cfg.ConnectionURL)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}
	return &Repository{
		db: dbs,
	}, nil
}

func (r *Repository) CreateUser(ctx context.Context, username, password string) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO users (username, password, created_at) VALUES ($1,$2)`, username, password)
	if err != nil {
		return fmt.Errorf("db exec context: %w", err)
	}

	return nil
}
