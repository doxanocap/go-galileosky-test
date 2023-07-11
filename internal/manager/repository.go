package manager

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryManager struct {
	pool *pgxpool.Pool
}

func InitRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		pool: pool,
	}
}
