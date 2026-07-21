package dbtx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryxContext(context.Context, string, ...interface{}) (*sqlx.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	QueryRowxContext(context.Context, string, ...interface{}) *sqlx.Row
}
