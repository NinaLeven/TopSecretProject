package storage

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/NinaLeven/TopSecretProject/internal/projectmanager"
)

type Factory interface {
	NewRepository() projectmanager.Storage
	RunInTransaction(ctx context.Context, f func(context.Context, projectmanager.Storage) error) error
}

type factory struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewFactory(db *sqlx.DB) Factory {

}

func (f *factory) NewRepository() projectmanager.Storage {
	return &Repository{db: f.db}
}

func (f *factory) RunInTransaction(ctx context.Context, fun func(context.Context, projectmanager.Storage) error) error {
	tx, err := f.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %w", err)
	}

	err = fun(ctx, &Repository{db: tx})
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			f.log.WithError(rerr).Error("unable to rollback transaction: %s", rerr)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}

	return nil
}

type sqlxDB interface {
	squirrel.ExecerContext
	squirrel.QueryerContext
}

type Repository struct {
	db sqlxDB
}
