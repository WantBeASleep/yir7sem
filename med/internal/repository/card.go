package repository

import (
	"fmt"

	"yirv2/med/internal/repository/entity"
	"yirv2/pkg/daolib"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const cardTable = "card"

type CardQuery interface {
	InsertCard(card entity.Card) error
	GetCardByPK(doctorID uuid.UUID, patientID uuid.UUID) (entity.Card, error)
	CheckCardExists(doctorID uuid.UUID, patientID uuid.UUID) (bool, error)
	UpdateCard(card entity.Card) (int64, error)
}

type cardQuery struct {
	*daolib.BaseQuery
}

func (q *cardQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *cardQuery) InsertCard(card entity.Card) error {
	query := q.QueryBuilder().
		Insert(cardTable).
		Columns(
			"doctor_id",
			"patient_id",
			"diagnosis",
		).
		Values(
			card.DoctorID,
			card.PatientID,
			card.Diagnosis,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert card: %w", err)
	}

	return nil
}

func (q *cardQuery) GetCardByPK(doctorID uuid.UUID, patientID uuid.UUID) (entity.Card, error) {
	query := q.QueryBuilder().
		Select(
			"doctor_id",
			"patient_id",
			"diagnosis",
		).
		From(cardTable).
		Where(sq.Eq{
			"doctor_id":  doctorID,
			"patient_id": patientID,
		})

	var card entity.Card
	if err := q.Runner().Getx(q.Context(), &card, query); err != nil {
		return entity.Card{}, fmt.Errorf("get card: %w", err)
	}

	return card, nil
}

func (q *cardQuery) CheckCardExists(doctorID uuid.UUID, patientID uuid.UUID) (bool, error) {
	query := q.QueryBuilder().
		Select("1").
		Prefix("SELECT EXISTS (").
		From(cardTable).
		Where(sq.Eq{
			"doctor_id":  doctorID,
			"patient_id": patientID,
		}).
		Suffix(")")

	var exists bool
	if err := q.Runner().Getx(q.Context(), &exists, query); err != nil {
		return false, fmt.Errorf("check card exists by pk: %w", err)
	}

	return exists, nil
}

func (q *cardQuery) UpdateCard(card entity.Card) (int64, error) {
	query := q.QueryBuilder().
		Update(cardTable).
		SetMap(sq.Eq{
			"diagnosis": card.Diagnosis,
		}).
		Where(sq.Eq{
			"doctor_id":  card.DoctorID,
			"patient_id": card.PatientID,
		})

	res, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, fmt.Errorf("update card: %w", err)
	}

	return res.RowsAffected()
}
