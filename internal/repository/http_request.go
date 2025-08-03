package repository

import (
	"context"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

type HTTPRequestRepository struct {
	dbPool *pgxpool.Pool
}

func NewHTTPRequestRepository(dbPool *pgxpool.Pool) *HTTPRequestRepository {
	return &HTTPRequestRepository{dbPool: dbPool}
}

func (r *HTTPRequestRepository) CreateHTTPRequest(ctx context.Context) error {
	return nil
}
