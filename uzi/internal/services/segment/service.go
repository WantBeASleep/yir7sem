package segment

import (
	"context"
	"errors"
	"fmt"

	"uzi/internal/domain"
	"uzi/internal/repository"
	"uzi/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateSegment(ctx context.Context, segment domain.Segment) (uuid.UUID, error)
	UpdateSegment(ctx context.Context, id uuid.UUID, update UpdateSegment) (domain.Segment, error)
	DeleteSegment(ctx context.Context, id uuid.UUID) error
}

type service struct {
	dao repository.DAO
}

func New(
	dao repository.DAO,
) Service {
	return &service{
		dao: dao,
	}
}

func (s *service) CreateSegment(ctx context.Context, segment domain.Segment) (uuid.UUID, error) {
	segment.Id = uuid.New()
	if err := s.dao.NewSegmentQuery(ctx).InsertSegment(entity.Segment{}.FromDomain(segment)); err != nil {
		return uuid.Nil, fmt.Errorf("insert segment: %w", err)
	}

	return segment.Id, nil
}

func (s *service) UpdateSegment(ctx context.Context, id uuid.UUID, update UpdateSegment) (domain.Segment, error) {
	segmentQuery := s.dao.NewSegmentQuery(ctx)
	segmentDB, err := segmentQuery.GetSegmentByPK(id)
	if err != nil {
		return domain.Segment{}, fmt.Errorf("get segment by pk: %w", err)
	}
	segment := segmentDB.ToDomain()
	update.Update(&segment)

	_, err = segmentQuery.UpdateSegment(entity.Segment{}.FromDomain(segment))
	if err != nil {
		return domain.Segment{}, fmt.Errorf("update segment: %w", err)
	}

	return segment, nil
}

func (s *service) DeleteSegment(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	segmentQuery := s.dao.NewSegmentQuery(ctx)

	segment, err := segmentQuery.GetSegmentByPK(id)
	if err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("get segment by pk: %w", errors.Join(err, rollbackErr))
	}

	if err := segmentQuery.DeleteSegmentByPK(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete segment: %w", errors.Join(err, rollbackErr))
	}

	remainingSegments, err := segmentQuery.GetSegmentsByNodeID(segment.NodeID)
	if err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("get segment by node_id: %w", errors.Join(err, rollbackErr))
	}

	// у node не осталось сегментов, удаляем
	if len(remainingSegments) == 0 {
		if err := s.dao.NewNodeQuery(ctx).DeleteNodeByPK(segment.NodeID); err != nil {
			rollbackErr := s.dao.RollbackTx(ctx)
			return fmt.Errorf("delete node by id: %w", errors.Join(err, rollbackErr))
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
