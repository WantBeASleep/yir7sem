package repository

import (
	"context"

	"github.com/WantBeASleep/goooool/daolib"

	"github.com/jmoiron/sqlx"
)

type DAO interface {
	daolib.DAO
	NewDoctorQuery(ctx context.Context) DoctorQuery
	NewPatientQuery(ctx context.Context) PatientQuery
	NewCardQuery(ctx context.Context) CardQuery
}

type dao struct {
	daolib.DAO
}

func NewRepository(psql *sqlx.DB) DAO {
	return &dao{DAO: daolib.NewDao(psql)}
}

func (d *dao) NewDoctorQuery(ctx context.Context) DoctorQuery {
	doctorQuery := &doctorQuery{}
	d.NewRepo(ctx, doctorQuery)

	return doctorQuery
}

func (d *dao) NewPatientQuery(ctx context.Context) PatientQuery {
	patientQuery := &patientQuery{}
	d.NewRepo(ctx, patientQuery)

	return patientQuery
}

func (d *dao) NewCardQuery(ctx context.Context) CardQuery {
	cardQuery := &cardQuery{}
	d.NewRepo(ctx, cardQuery)

	return cardQuery
}
