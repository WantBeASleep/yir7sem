package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const imageTable = "image"

type ImageQuery interface {
	InsertImages(images []entity.Image) error
	GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error)
}

type imageQuery struct {
	*daolib.BaseQuery
}

func (q *imageQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *imageQuery) InsertImages(images []entity.Image) error {
	query := q.QueryBuilder().
		Insert(imageTable).
		Columns(
			"id",
			"uzi_id",
			"page",
		)

	for _, v := range images {
		query = query.Values(
			v.Id,
			v.UziID,
			v.Page,
		)
	}

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return err
}

func (q *imageQuery) GetImagesByUziID(uziID uuid.UUID) ([]entity.Image, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"page",
		).
		From(imageTable).
		Where(sq.Eq{
			"uzi_id": uziID,
		})

	var images []entity.Image
	if err := q.Runner().Selectx(q.Context(), &images, query); err != nil {
		return nil, err
	}

	return images, nil
}
