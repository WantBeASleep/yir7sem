package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"med/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const patientTable = "patient"

type PatientQuery interface {
	InsertPatient(patient entity.Patient) error
	GetPatientByPK(id uuid.UUID) (entity.Patient, error)
	GetPatientsByDoctorID(id uuid.UUID) ([]entity.Patient, error)
	UpdatePatient(patient entity.Patient) (int64, error)
}

type patientQuery struct {
	*daolib.BaseQuery
}

func (q *patientQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *patientQuery) InsertPatient(patient entity.Patient) error {
	query := q.QueryBuilder().
		Insert(patientTable).
		Columns(
			"id",
			"fullname",
			"email",
			"policy",
			"active",
			"malignancy",
			"last_uzi_date",
		).
		Values(
			patient.Id,
			patient.FullName,
			patient.Email,
			patient.Policy,
			patient.Active,
			patient.Malignancy,
			patient.LastUziDate,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *patientQuery) GetPatientByPK(id uuid.UUID) (entity.Patient, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"fullname",
			"email",
			"policy",
			"active",
			"malignancy",
			"last_uzi_date",
		).
		From(patientTable).
		Where(sq.Eq{
			"id": id,
		})

	var patient entity.Patient
	if err := q.Runner().Getx(q.Context(), &patient, query); err != nil {
		return entity.Patient{}, err
	}

	return patient, nil
}

func (q *patientQuery) GetPatientsByDoctorID(id uuid.UUID) ([]entity.Patient, error) {
	query := q.QueryBuilder().
		Select(
			"patient.id",
			"patient.fullname",
			"patient.email",
			"patient.policy",
			"patient.active",
			"patient.malignancy",
			"patient.last_uzi_date",
		).
		From(patientTable).
		InnerJoin("card ON card.patient_id = patient.id").
		Where(sq.Eq{
			"card.doctor_id": id,
		})

	var patient []entity.Patient
	if err := q.Runner().Selectx(q.Context(), &patient, query); err != nil {
		return nil, err
	}

	return patient, nil
}

func (q *patientQuery) UpdatePatient(patient entity.Patient) (int64, error) {
	query := q.QueryBuilder().
		Update(patientTable).
		SetMap(sq.Eq{
			"active":        patient.Active,
			"malignancy":    patient.Malignancy,
			"last_uzi_date": patient.LastUziDate,
		}).
		Where(sq.Eq{
			"id": patient.Id,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
