package storage

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type Factory interface {
	NewRepository() projectmanager.Storage
	RunInTransaction(ctx context.Context, f func(context.Context, projectmanager.Storage) error) error
}

type sqlxDB interface {
	sqlx.PreparerContext
	sqlx.QueryerContext
	sqlx.ExecerContext
}

type Repository struct {
	db sqlxDB
}
