package repository

import (
	"fmt"

	"pkg/daolib"
	"uzi/internal/repository/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const nodeTable = "node"

type NodeQuery interface {
	InsertNode(node entity.Node) error
	GetNodeByPK(id uuid.UUID) (entity.Node, error)
	GetNodesByImageID(id uuid.UUID) ([]entity.Node, error)
	UpdateNode(node entity.Node) (int64, error)
	DeleteNodeByPK(id uuid.UUID) error
}

type nodeQuery struct {
	*daolib.BaseQuery
}

func (q *nodeQuery) SetBaseQuery(baseQuery *daolib.BaseQuery) {
	q.BaseQuery = baseQuery
}

func (q *nodeQuery) InsertNode(node entity.Node) error {
	query := q.QueryBuilder().
		Insert(nodeTable).
		Columns(
			"id",
			"ai",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		Values(
			node.Id,
			node.Ai,
			node.Tirads23,
			node.Tirads4,
			node.Tirads5,
		)

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("insert node: %w", err)
	}

	return nil
}

func (q *nodeQuery) GetNodeByPK(id uuid.UUID) (entity.Node, error) {
	query := q.QueryBuilder().
		Select(
			"id",
			"ai",
			"tirads_23",
			"tirads_4",
			"tirads_5",
		).
		From(nodeTable).
		Where(sq.Eq{
			"id": id,
		})

	var node entity.Node
	if err := q.Runner().Getx(q.Context(), &node, query); err != nil {
		return entity.Node{}, fmt.Errorf("get node: %w", err)
	}

	return node, nil
}

func (q *nodeQuery) GetNodesByImageID(id uuid.UUID) ([]entity.Node, error) {
	query := q.QueryBuilder().
		Select(
			"node.id",
			"node.ai",
			"node.tirads_23",
			"node.tirads_4",
			"node.tirads_5",
		).
		From(nodeTable).
		InnerJoin("segment ON segment.node_id = node.id").
		InnerJoin("image ON image.id = segment.image_id").
		Where(sq.Eq{
			"image.id": id,
		})

	var uzi []entity.Node
	if err := q.Runner().Selectx(q.Context(), &uzi, query); err != nil {
		return nil, fmt.Errorf("get node by image_id: %w", err)
	}

	return uzi, nil
}

// TODO: упорядочнить Insert/Delete/Select/Update
func (q *nodeQuery) UpdateNode(node entity.Node) (int64, error) {
	query := q.QueryBuilder().
		Update(nodeTable).
		SetMap(sq.Eq{
			"tirads_23": node.Tirads23,
			"tirads_4":  node.Tirads4,
			"tirads_5":  node.Tirads5,
		}).
		Where(sq.Eq{
			"id": node.Id,
		})

	rows, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return 0, fmt.Errorf("update node: %w", err)
	}

	return rows.RowsAffected()
}

func (q *nodeQuery) DeleteNodeByPK(id uuid.UUID) error {
	query := q.QueryBuilder().
		Delete(nodeTable).
		Where(sq.Eq{
			"id": id,
		})

	_, err := q.Runner().Execx(q.Context(), query)
	if err != nil {
		return fmt.Errorf("delete node: %w", err)
	}

	return nil
}
