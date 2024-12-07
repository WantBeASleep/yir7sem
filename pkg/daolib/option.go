package daolib

import (
	"database/sql"
)

type TxOption func(*sql.TxOptions)

func WithIsolationLevel(lvl sql.IsolationLevel) TxOption {
	return TxOption(func(o *sql.TxOptions) {
		o.Isolation = lvl
	})
}

func WithReadOnly() TxOption {
	return TxOption(func(o *sql.TxOptions) {
		o.ReadOnly = true
	})
}
