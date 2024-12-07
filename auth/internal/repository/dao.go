package repository

import (
	"context"
	"pkg/daolib"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewUserQuery(ctx context.Context) UserQuery
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewUserQuery(ctx context.Context) UserQuery {
	userQuery := &userQuery{}
	d.NewRepo(ctx, userQuery)

	return userQuery
}
