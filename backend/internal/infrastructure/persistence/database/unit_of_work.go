package database

import (
	"context"
	"database/sql"

	repositoryiface "github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/repository"

	dbtx "github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/dbtx"

	repositories "github.com/arthurhzna/inventory-reservation-system/backend/internal/infrastructure/persistence/repository"

	"github.com/jmoiron/sqlx"
)

type unitOfWork struct {
	conn *sqlx.DB
	db   dbtx.DBTX
}

func NewUnitOfWork(
	db dbtx.DBTX,
) repositoryiface.UnitOfWork {
	return &unitOfWork{
		conn: db.(*sqlx.DB),
		db:   db,
	}
}

func (u *unitOfWork) WithTransaction(
	ctx context.Context,
	fn func(repositoryiface.UnitOfWork) error,
) error {

	tx, err := u.conn.BeginTxx(
		ctx,
		&sql.TxOptions{},
	)
	if err != nil {
		return err
	}

	transactionalUow := &unitOfWork{
		conn: u.conn,
		db:   tx,
	}

	err = fn(transactionalUow)
	if err != nil {

		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return err
		}

		return err
	}

	return tx.Commit()
}

func (u *unitOfWork) UserRepository() repositoryiface.UserRepository {
	return repositories.NewUserRepository(u.db)
}
