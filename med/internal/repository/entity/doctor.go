package entity

import (
	"database/sql"

	"yir/med/internal/domain"
	"yir/pkg/gtclib"

	"github.com/google/uuid"
)

type Doctor struct {
	Id       uuid.UUID      `db:"id"`
	FullName string         `db:"fullname"`
	Org      string         `db:"org"`
	Job      string         `db:"job"`
	Desc     sql.NullString `db:"desc"`
}

func (Doctor) FromDomain(d domain.Doctor) Doctor {
	return Doctor{
		Id:       d.Id,
		FullName: d.FullName,
		Org:      d.Org,
		Job:      d.Job,
		Desc:     gtclib.String.PointerToSql(d.Desc),
	}
}

func (d Doctor) ToDomain() domain.Doctor {
	return domain.Doctor{
		Id:       d.Id,
		FullName: d.FullName,
		Org:      d.Org,
		Job:      d.Job,
		Desc:     gtclib.String.SqlToPointer(d.Desc),
	}
}
