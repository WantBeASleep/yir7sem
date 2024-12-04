package gtclib

import "database/sql"

type _float64 struct{}

var Float64 _float64

func (_float64) PointerToSql(p *float64) sql.NullFloat64 {
	if p == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Valid: true, Float64: *p}
}

func (_float64) SqlToPointer(sq sql.NullFloat64) *float64 {
	if !sq.Valid {
		return nil
	}

	v := sq.Float64
	return &v
}
