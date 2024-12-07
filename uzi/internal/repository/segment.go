package repository

import (
	"github.com/WantBeASleep/goooool/daolib"

	"uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const segmentTable = "segment"

type SegmentQuery interface {
	InsertSegment(segment entity.Segment) error
	GetSegmentByPK(id uuid.UUID) (entity.Segment, error)
	GetSegmentsByNodeID(id uuid.UUID) ([]entity.Segment, error)
	GetSegmentsByImageID(id uuid.UUID) ([]entity.Segment, error)
	// GetUziIDBySegmentID(id uuid.UUID) (uuid.UUID, error)
	UpdateSegment(segment entity.Segment) (int64, error)
	DeleteSegmentByPK(id uuid.UUID) error
	DeleteSegmentByUziID(id uuid.UUID) (int64, error)
}

type segmentQuery struct {
	*daolib.BaseQuery
}

func (q *segmentQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *segmentQuery) InsertSegment(segment entity.Segment) error {
	query := q.QueryBuilder().
		Insert(segmentTable).
		Columns(
			"id",
			"node_id",
			"image_id",
			"contor",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		Values(
			segment.Id,
			segment.NodeID,
			segment.ImageID,
			segment.Contor,
			segment.Tirads23,
			segment.Tirads4,
			segment.Tirads5,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *segmentQuery) GetSegmentByPK(id uuid.UUID) (entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"node_id",
			"image_id",
			"contor",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		From(segmentTable).
		Where(sq.Eq{
			"id": id,
		})

	var segments entity.Segment
	if err := q.Runner().Getx(q.Context(), &segments, query); err != nil {
		return entity.Segment{}, err
	}

	return segments, nil
}

func (q *segmentQuery) GetSegmentsByNodeID(id uuid.UUID) ([]entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"node_id",
			"image_id",
			"contor",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		From(segmentTable).
		Where(sq.Eq{
			"node_id": id,
		})

	var segments []entity.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, err
	}

	return segments, nil
}

func (q *segmentQuery) GetSegmentsByImageID(id uuid.UUID) ([]entity.Segment, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"node_id",
			"image_id",
			"contor",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		From(segmentTable).
		Where(sq.Eq{
			"image_id": id,
		})

	var segments []entity.Segment
	if err := q.Runner().Selectx(q.Context(), &segments, query); err != nil {
		return nil, err
	}

	return segments, nil
}

// func (q *segmentQuery) GetUziIDBySegmentID(id uuid.UUID) (uuid.UUID, error) {
// 	query := q.QueryBuilder().
// 		Select(
// 			"uzi.id",
// 		).
// 		From(segmentTable).
// 		InnerJoin("image ON image.id = segment.image_id").
// 		InnerJoin("uzi ON image.uzi_id = uzi.id").
// 		Where(sq.Eq{
// 			"segment.id": id,
// 		})

// 	var uziID uuid.UUID
// 	if err := q.Runner().Getx(q.Context(), &uziID, query); err != nil {
// 		return uuid.Nil, err
// 	}

// 	return uziID, nil
// }

func (q *segmentQuery) UpdateSegment(segment entity.Segment) (int64, error) {
	query := q.QueryBuilder().
		Update(segmentTable).
		SetMap(sq.Eq{
			"tirads_23": segment.Tirads23,
			"tirads_4":  segment.Tirads4,
			"tirads_5":  segment.Tirads5,
		}).
		Where(sq.Eq{
			"id": segment.Id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}

func (q *segmentQuery) DeleteSegmentByPK(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(segmentTable).
		Where(sq.Eq{
			"id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return err
	}

	return nil
}

func (q *segmentQuery) DeleteSegmentByUziID(id uuid.UUID) (int64, error) {
	query := q.QueryBuilder().
		Delete(segmentTable).
		Where(sq.Eq{
			"node_id": id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, err
	}

	return rows.RowsAffected()
}
