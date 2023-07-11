package manager

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"sync"
	"todo/internal/manager/interfaces/repository"
	repository2 "todo/internal/repository"
	"todo/pkg/logger"
)

type RepositoryManager struct {
	pool *pgxpool.Pool

	todoItem       repository.ITodoItemRepository
	todoItemRunner sync.Once
}

func InitRepositoryManager(pool *pgxpool.Pool) *RepositoryManager {
	return &RepositoryManager{
		pool: pool,
	}
}

func (rm *RepositoryManager) TodoItem() repository.ITodoItemRepository {
	rm.todoItemRunner.Do(func() {
		rm.todoItem = repository2.InitTodoItemsRepository(rm.pool, logger.Log.Named("[REPOSITORY][TODO_ITEM]"))
	})
	return rm.todoItem
}
