package entity

import (
	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Segment struct {
	Id       uuid.UUID `db:"id"`
	ImageID  uuid.UUID `db:"image_id"`
	NodeID   uuid.UUID `db:"node_id"`
	Contor   string    `db:"contor"`
	Tirads23 float64   `db:"tirads_23"`
	Tirads4  float64   `db:"tirads_4"`
	Tirads5  float64   `db:"tirads_5"`
}

func (Segment) FromDomain(d domain.Segment) Segment {
	return Segment{
		Id:       d.Id,
		ImageID:  d.ImageID,
		NodeID:   d.NodeID,
		Contor:   d.Contor,
		Tirads23: d.Tirads23,
		Tirads4:  d.Tirads4,
		Tirads5:  d.Tirads5,
	}
}

func (d Segment) ToDomain() domain.Segment {
	return domain.Segment{
		Id:       d.Id,
		ImageID:  d.ImageID,
		NodeID:   d.NodeID,
		Contor:   d.Contor,
		Tirads23: d.Tirads23,
		Tirads4:  d.Tirads4,
		Tirads5:  d.Tirads5,
	}
}

func (Segment) SliceToDomain(slice []Segment) []domain.Segment {
	res := make([]domain.Segment, 0, len(slice))
	for _, v := range slice {
		res = append(res, v.ToDomain())
	}
	return res
}
