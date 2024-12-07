package entity

import (
	"database/sql"

	"github.com/WantBeASleep/goooool/gtclib"

	"med/internal/domain"

	"github.com/google/uuid"
)

type Card struct {
	DoctorID  uuid.UUID      `db:"doctor_id"`
	PatientID uuid.UUID      `db:"patient_id"`
	Diagnosis sql.NullString `db:"diagnosis"`
}

func (Card) FromDomain(p domain.Card) Card {
	return Card{
		DoctorID:  p.DoctorID,
		PatientID: p.PatientID,
		Diagnosis: gtclib.String.PointerToSql(p.Diagnosis),
	}
}

func (p Card) ToDomain() domain.Card {
	return domain.Card{
		DoctorID:  p.DoctorID,
		PatientID: p.PatientID,
		Diagnosis: gtclib.String.SqlToPointer(p.Diagnosis),
	}
}
