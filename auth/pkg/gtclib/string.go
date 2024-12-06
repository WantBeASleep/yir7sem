package gtclib

import "database/sql"

type _string struct{}

var String _string

func (_string) ValueToPointer(p string) *string {
	if p == "" {
		return nil
	}

	return &p
}

func (_string) PointerToSql(p *string) sql.NullString {
	if p == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *p, Valid: true}
}

func (_string) SqlToPointer(sq sql.NullString) *string {
	if !sq.Valid {
		return nil
	}

	v := sq.String
	return &v
}
