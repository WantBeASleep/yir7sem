package entity

import (
	"database/sql"

	"github.com/WantBeASleep/goooool/gtclib"

	"med/internal/domain"

	"github.com/google/uuid"
)

type Patient struct {
	Id          uuid.UUID    `db:"id"`
	FullName    string       `db:"fullname"`
	Email       string       `db:"email"`
	Policy      string       `db:"policy"`
	Active      bool         `db:"active"`
	Malignancy  bool         `db:"malignancy"`
	LastUziDate sql.NullTime `db:"last_uzi_date"`
}

// TODO: пройтись по таблице перевести NULLSQL на этот тип
func (Patient) FromDomain(p domain.Patient) Patient {
	return Patient{
		Id:          p.Id,
		FullName:    p.FullName,
		Email:       p.Email,
		Policy:      p.Policy,
		Active:      p.Active,
		Malignancy:  p.Malignancy,
		LastUziDate: gtclib.Time.PointerToSql(p.LastUziDate),
	}
}

func (p Patient) ToDomain() domain.Patient {
	return domain.Patient{
		Id:          p.Id,
		FullName:    p.FullName,
		Email:       p.Email,
		Policy:      p.Policy,
		Active:      p.Active,
		Malignancy:  p.Malignancy,
		LastUziDate: gtclib.Time.SqlToPointer(p.LastUziDate),
	}
}
