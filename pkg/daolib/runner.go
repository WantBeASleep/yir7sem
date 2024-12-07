package daolib

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type Builder interface {
	ToSql() (sql string, args []any, err error)
}

type Runner interface {
	Execx(ctx context.Context, builder Builder) (sql.Result, error)
	Getx(ctx context.Context, dest interface{}, builder Builder) error
	Selectx(ctx context.Context, dest interface{}, builder Builder) error
	QueryRow(ctx context.Context, builder Builder) (*sqlx.Row, error)
}

type Tx struct {
	tx *sqlx.Tx
}

func NewTx(tx *sqlx.Tx) *Tx {
	return &Tx{tx: tx}
}

func (tx *Tx) Execx(ctx context.Context, builder Builder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "tx exec", slog.String("query", query), slog.Any("args", args))
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *Tx) Getx(ctx context.Context, dest interface{}, builder Builder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "tx getx", slog.String("query", query), slog.Any("args", args))
	return tx.tx.GetContext(ctx, dest, query, args...)
}

func (tx *Tx) Selectx(ctx context.Context, dest interface{}, builder Builder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "tx selectx", slog.String("query", query), slog.Any("args", args))
	return tx.tx.SelectContext(ctx, dest, query, args...)
}

func (tx *Tx) QueryRow(ctx context.Context, builder Builder) (*sqlx.Row, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "tx query row", slog.String("query", query), slog.Any("args", args))
	return tx.tx.QueryRowxContext(ctx, query, args...), nil
}

type DB struct {
	db *sqlx.DB
}

func NewDB(db *sqlx.DB) *DB {
	return &DB{db: db}
}

func (db *DB) Execx(ctx context.Context, builder Builder) (sql.Result, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "no tx exec", slog.String("query", query), slog.Any("args", args))
	return db.db.ExecContext(ctx, query, args...)
}

func (db *DB) Getx(ctx context.Context, dest interface{}, builder Builder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "no tx getx", slog.String("query", query), slog.Any("args", args))
	return db.db.GetContext(ctx, dest, query, args...)
}

func (db *DB) Selectx(ctx context.Context, dest interface{}, builder Builder) error {
	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "no tx selectx", slog.String("query", query), slog.Any("args", args))
	return db.db.SelectContext(ctx, dest, query, args...)
}

func (db *DB) QueryRow(ctx context.Context, builder Builder) (*sqlx.Row, error) {
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build sql req: %w", err)
	}

	slog.DebugContext(ctx, "no tx query row", slog.String("query", query), slog.Any("args", args))
	return db.db.QueryRowxContext(ctx, query, args...), nil
}
