package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Node struct {
	Id       uuid.UUID `db:"id"`
	Ai       bool      `db:"ai"`
	UziID    uuid.UUID `db:"uzi_id"`
	Tirads23 float64   `db:"tirads_23"`
	Tirads4  float64   `db:"tirads_4"`
	Tirads5  float64   `db:"tirads_5"`
}

func (Node) FromDomain(d domain.Node) Node {
	return Node{
		Id:       d.Id,
		Ai:       d.Ai,
		UziID:    d.UziID,
		Tirads23: d.Tirads23,
		Tirads4:  d.Tirads4,
		Tirads5:  d.Tirads5,
	}
}

func (d Node) ToDomain() domain.Node {
	return domain.Node{
		Id:       d.Id,
		Ai:       d.Ai,
		UziID:    d.UziID,
		Tirads23: d.Tirads23,
		Tirads4:  d.Tirads4,
		Tirads5:  d.Tirads5,
	}
}

func (Node) SliceToDomain(slice []Node) []domain.Node {
	res := make([]domain.Node, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
