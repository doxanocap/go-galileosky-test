package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"time"
	"todo/internal/cns"
	"todo/internal/model"
)

type TodoItemsRepository struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	builder sq.StatementBuilderType
}

func InitTodoItemsRepository(pool *pgxpool.Pool, log *zap.Logger) *TodoItemsRepository {
	return &TodoItemsRepository{
		pool:    pool,
		log:     log,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repo *TodoItemsRepository) Create(ctx context.Context, request *model.TodoItem) (result *[]model.TodoItem, err error) {
	log := repo.log.Named("Create").With(
		zap.String(cns.TodoItemsTable, request.Title),
		zap.String(cns.TodoItemsDesc, request.Description))

	result = &[]model.TodoItem{}

	query := repo.builder.
		Insert(cns.TodoItemsTable).
		Columns(
			cns.TodoItemsTitle,
			cns.TodoItemsDesc,
		).
		Values(
			request.Title,
			request.Description,
		).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *TodoItemsRepository) GetAll(ctx context.Context) (result *[]model.TodoItem, err error) {
	log := repo.log.Named("GetAll")

	result = &[]model.TodoItem{}

	query := repo.builder.
		Select("*").
		From(cns.TodoItemsTable)

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *TodoItemsRepository) GetByID(ctx context.Context, ID int64) (result *[]model.TodoItem, err error) {
	log := repo.log.Named("GetByID").With(
		zap.Int64(cns.TodoItemsID, ID))

	result = &[]model.TodoItem{}

	query := repo.builder.
		Select("*").
		From(cns.TodoItemsTable).
		Where(sq.And{
			sq.Eq{cns.TodoItemsID: ID},
			sq.Eq{cns.TodoItemDeletedAt: nil},
		})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *TodoItemsRepository) GetByTitle(ctx context.Context, title string) (result *[]model.TodoItem, err error) {
	log := repo.log.Named("GetByTitle").With(
		zap.String(cns.TodoItemsTitle, title))

	result = &[]model.TodoItem{}

	query := repo.builder.
		Select("*").
		From(cns.TodoItemsTable).
		Where(sq.And{
			sq.Eq{cns.TodoItemsTitle: title},
			sq.Eq{cns.TodoItemDeletedAt: nil},
		})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *TodoItemsRepository) UpdateByID(ctx context.Context, request *model.TodoItem) (err error) {
	log := repo.log.Named("UpdateByID").With(
		zap.Int64(cns.TodoItemsID, request.Id),
		zap.String(cns.TodoItemsTitle, request.Title),
		zap.String(cns.TodoItemsDesc, request.Description),
		zap.Bool(cns.TodoItemDone, request.Done))

	query := repo.builder.
		Update(cns.TodoItemsTable).
		SetMap(map[string]interface{}{
			cns.TodoItemsTitle: request.Title,
			cns.TodoItemsDesc:  request.Description,
			cns.TodoItemDone:   request.Done,
		}).
		Where(sq.And{
			sq.Eq{cns.TodoItemsID: request.Id},
			sq.Eq{cns.TodoItemDeletedAt: nil},
		})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	_, err = repo.pool.Exec(ctx, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *TodoItemsRepository) DeleteByID(ctx context.Context, ID int64) (err error) {
	log := repo.log.Named("DeleteByID").With(
		zap.Int64(cns.TodoItemsID, ID))

	query := repo.builder.
		Update(cns.TodoItemsTable).
		SetMap(map[string]interface{}{
			cns.TodoItemDeletedAt: time.Now(),
		}).
		Where(sq.Eq{cns.TodoItemsID: ID}).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	_, err = repo.pool.Exec(ctx, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}
