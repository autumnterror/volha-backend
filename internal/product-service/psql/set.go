package psql

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrUnknownType = errors.New("unknown type")
	ErrInvalidType = errors.New("bad type of obj")
)

type SqlRepo interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Driver struct {
	Driver SqlRepo
}
