package daolib

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

type BaseQuerySetter interface {
	SetBaseQuery(querier *BaseQuery)
}

type BaseQuery struct {
	ctx    context.Context
	runner Runner
}

func (q *BaseQuery) Context() context.Context {
	return q.ctx
}

func (q *BaseQuery) Runner() Runner {
	return q.runner
}

func (*BaseQuery) QueryBuilder() sq.StatementBuilderType {
	return sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
