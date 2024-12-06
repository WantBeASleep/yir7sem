package repository

import (
	"fmt"

	"uzi/internal/repository/entity"
	"uzi/pkg/daolib"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const segmentTable = "segment"

type SegmentQuery interface {
	InsertSegment(segment entity.Segment) error
	GetSegmentByPK(id uuid.UUID) (entity.Segment, error)
	GetSegmentsByNodeID(id uuid.UUID) ([]entity.Segment, error)
	GetSegmentsByImageID(id uuid.UUID) ([]entity.Segment, error)
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
		return fmt.Errorf("insert segment: %w", err)
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
		return entity.Segment{}, fmt.Errorf("get segments by pk: %w", err)
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
		return nil, fmt.Errorf("get segments by uzi_id: %w", err)
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
		return nil, fmt.Errorf("get segments by image_id: %w", err)
	}

	return segments, nil
}

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
		return 0, fmt.Errorf("update segment: %w", err)
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
		return fmt.Errorf("delete segment: %w", err)
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
		return 0, fmt.Errorf("delete segment by uzi_id: %w", err)
	}

	return rows.RowsAffected()
}
