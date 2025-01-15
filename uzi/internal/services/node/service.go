package node

import (
	"context"
	"errors"
	"fmt"

	// "uzi/internal/adapters"

	"uzi/internal/domain"
	"uzi/internal/repository"
	"uzi/internal/repository/entity"

	"github.com/google/uuid"
)

type Service interface {
	CreateNode(ctx context.Context, node domain.Node, segments []domain.Segment) (uuid.UUID, error)
	InsertAiNodeWithSegments(ctx context.Context, nodes []domain.Node, segments []domain.Segment) error
	GetAllNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, error)
	UpdateNode(ctx context.Context, id uuid.UUID, update UpdateNode) (domain.Node, error)
	DeleteNode(ctx context.Context, id uuid.UUID) error
}

type service struct {
	dao repository.DAO
	// adapter adapters.Adapter
}

func New(
	dao repository.DAO,
	// adapter adapters.Adapter,
) Service {
	return &service{
		dao: dao,
		// adapter: adapter,
	}
}

// TODO: rollback через defer
func (s *service) CreateNode(ctx context.Context, node domain.Node, segments []domain.Segment) (uuid.UUID, error) {
	node.Id = uuid.New()
	// ai create через брокер
	node.Ai = false

	for i := range segments {
		segments[i].Id = uuid.New()
		segments[i].NodeID = node.Id
	}

	if err := s.insertNodesWithSegments(ctx, []domain.Node{node}, segments); err != nil {
		return uuid.Nil, fmt.Errorf("insert node with segments: %w", err)
	}

	return node.Id, nil
}

func (s *service) InsertAiNodeWithSegments(ctx context.Context, nodes []domain.Node, segments []domain.Segment) error {
	for i := range nodes {
		nodes[i].Ai = true
	}

	if err := s.insertNodesWithSegments(ctx, nodes, segments); err != nil {
		return fmt.Errorf("insert node with segments: %w", err)
	}

	// TODO: подумать тут мб добавить еще поля в node и segments, что бы не делать этот join
	// if err := s.adapter.BrokerAdapter.SendUziComplete(&uzicompletepb.UziComplete{
	// 	// UziId: ,
	// }); err != nil {
	// 	return fmt.Errorf("send to uzisplitted topic: %w", err)
	// }

	return nil
}

func (s *service) insertNodesWithSegments(ctx context.Context, nodes []domain.Node, segments []domain.Segment) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	nodeQuery := s.dao.NewNodeQuery(ctx)
	segmentQuery := s.dao.NewSegmentQuery(ctx)

	for _, v := range nodes {
		if err := nodeQuery.InsertNode(entity.Node{}.FromDomain(v)); err != nil {
			// TODO: rollback нормально ошибку оформить
			rollbackErr := s.dao.RollbackTx(ctx)
			return fmt.Errorf("insert node: %w", errors.Join(err, rollbackErr))
		}
	}

	for _, v := range segments {
		if err := segmentQuery.InsertSegment(entity.Segment{}.FromDomain(v)); err != nil {
			rollbackErr := s.dao.RollbackTx(ctx)
			return fmt.Errorf("insert segment: %w", errors.Join(err, rollbackErr))
		}
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (s *service) GetAllNodes(ctx context.Context, id uuid.UUID) ([]domain.Node, error) {
	nodesDB, err := s.dao.NewNodeQuery(ctx).GetNodesByUziID(id)
	if err != nil {
		return nil, err
	}

	return entity.Node{}.SliceToDomain(nodesDB), err
}

func (s *service) UpdateNode(ctx context.Context, id uuid.UUID, update UpdateNode) (domain.Node, error) {
	nodeQuery := s.dao.NewNodeQuery(ctx)

	nodeDB, err := nodeQuery.GetNodeByPK(id)
	if err != nil {
		return domain.Node{}, fmt.Errorf("get node: %w", err)
	}
	node := nodeDB.ToDomain()
	update.Update(&node)

	_, err = nodeQuery.UpdateNode(entity.Node{}.FromDomain(node))
	if err != nil {
		return domain.Node{}, fmt.Errorf("update node: %w", err)
	}

	return node, nil
}

func (s *service) DeleteNode(ctx context.Context, id uuid.UUID) error {
	ctx, err := s.dao.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	if _, err := s.dao.NewSegmentQuery(ctx).DeleteSegmentByUziID(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete node segments: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.NewNodeQuery(ctx).DeleteNodeByPK(id); err != nil {
		rollbackErr := s.dao.RollbackTx(ctx)
		return fmt.Errorf("delete node: %w", errors.Join(err, rollbackErr))
	}

	if err := s.dao.CommitTx(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
