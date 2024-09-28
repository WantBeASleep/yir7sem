package entity

import (
	"github.com/google/uuid"
)

type DBFormation struct {
	Id        uuid.UUID
	Segments  map[uuid.UUID]HttpSegment // сегменты узла
	AvgTirads Tirads
	Ai        bool
}
