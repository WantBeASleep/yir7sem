package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"med/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const doctorTable = "doctor"

type DoctorQuery interface {
	InsertDoctor(doctor entity.Doctor) error
	GetDoctorByPK(id uuid.UUID) (entity.Doctor, error)
	UpdateDoctor(doctor entity.Doctor) (int64, error)
}

type doctorQuery struct {
	*daolib.BaseQuery
}

func (q *doctorQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *doctorQuery) InsertDoctor(doctor entity.Doctor) error {
	query := q.QueryBuilder().
		Insert(doctorTable).
		Columns(
			"id",
			"fullname",
			"org",
			"job",
			"\"desc\"",
		).
		Values(
			doctor.Id,
			doctor.FullName,
			doctor.Org,
			doctor.Job,
			doctor.Desc,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *doctorQuery) GetDoctorByPK(id uuid.UUID) (entity.Doctor, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"fullname",
			"org",
			"job",
			"\"desc\"",
		).
		From(doctorTable).
		Where(sq.Eq{
			"id": id,
		})

	var doctor entity.Doctor
	if err := q.Runner().Getx(q.Context(), &doctor, query); err != nil {
		return entity.Doctor{}, err
	}

	return doctor, nil
}

func (q *doctorQuery) UpdateDoctor(doctor entity.Doctor) (int64, error) {
	query := q.QueryBuilder().
		Update(doctorTable).
		SetMap(sq.Eq{
			"org":      doctor.Org,
			"job":      doctor.Job,
			"\"desc\"": doctor.Desc,
		}).
		Where(sq.Eq{
			"id": doctor.Id,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
