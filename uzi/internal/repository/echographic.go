package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const echographicTable = "echographic"

type EchographicQuery interface {
	InsertEchographic(echographic entity.Echographic) error
	GetEchographicByPK(id uuid.UUID) (entity.Echographic, error)
	UpdateEchographic(echographic entity.Echographic) (int64, error)
}

type echographicQuery struct {
	*daolib.BaseQuery
}

func (q *echographicQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *echographicQuery) InsertEchographic(echographic entity.Echographic) error {
	query := q.QueryBuilder().
		Insert(echographicTable).
		Columns(
			"id",
			"contors",
			"left_lobe_length",
			"left_lobe_width",
			"left_lobe_thick",
			"left_lobe_volum",
			"right_lobe_length",
			"right_lobe_width",
			"right_lobe_thick",
			"right_lobe_volum",
			"gland_volum",
			"isthmus",
			"struct",
			"echogenicity",
			"regional_lymph",
			"vascularization",
			"location",
			"additional",
			"conclusion",
		).
		Values(
			echographic.Id,
			echographic.Contors,
			echographic.LeftLobeLength,
			echographic.LeftLobeWidth,
			echographic.LeftLobeThick,
			echographic.LeftLobeVolum,
			echographic.RightLobeLength,
			echographic.RightLobeWidth,
			echographic.RightLobeThick,
			echographic.RightLobeVolum,
			echographic.GlandVolum,
			echographic.Isthmus,
			echographic.Struct,
			echographic.Echogenicity,
			echographic.RegionalLymph,
			echographic.Vascularization,
			echographic.Location,
			echographic.Additional,
			echographic.Conclusion,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *echographicQuery) GetEchographicByPK(id uuid.UUID) (entity.Echographic, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"contors",
			"left_lobe_length",
			"left_lobe_width",
			"left_lobe_thick",
			"left_lobe_volum",
			"right_lobe_length",
			"right_lobe_width",
			"right_lobe_thick",
			"right_lobe_volum",
			"gland_volum",
			"isthmus",
			"struct",
			"echogenicity",
			"regional_lymph",
			"vascularization",
			"location",
			"additional",
			"conclusion",
		).
		From(echographicTable).
		Where(sq.Eq{
			"id": id,
		})

	var echographic entity.Echographic
	if err := q.Runner().Getx(q.Context(), &echographic, query); err != nil {
		return entity.Echographic{}, err
	}

	return echographic, nil
}

func (q *echographicQuery) UpdateEchographic(echographic entity.Echographic) (int64, error) {
	query := q.QueryBuilder().
		Update(echographicTable).
		SetMap(sq.Eq{
			"contors":           echographic.Contors,
			"left_lobe_length":  echographic.LeftLobeLength,
			"left_lobe_width":   echographic.LeftLobeWidth,
			"left_lobe_thick":   echographic.LeftLobeThick,
			"left_lobe_volum":   echographic.LeftLobeVolum,
			"right_lobe_length": echographic.RightLobeLength,
			"right_lobe_width":  echographic.RightLobeWidth,
			"right_lobe_thick":  echographic.RightLobeThick,
			"right_lobe_volum":  echographic.RightLobeVolum,
			"gland_volum":       echographic.GlandVolum,
			"isthmus":           echographic.Isthmus,
			"struct":            echographic.Struct,
			"echogenicity":      echographic.Echogenicity,
			"regional_lymph":    echographic.RegionalLymph,
			"vascularization":   echographic.Vascularization,
			"location":          echographic.Location,
			"additional":        echographic.Additional,
			"conclusion":        echographic.Conclusion,
		}).
		Where(sq.Eq{
			"id": echographic.Id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}
