package entity

import (
	"github.com/google/uuid"
)

type HttpImage struct {
	Id   uuid.UUID
	Url  string
	Page int
}

type HttpSegment struct {
	ContorUrl string
	Tirads Tirads
}

type HttpFormation struct {
	Id uuid.UUID
	Segments map[uuid.UUID]HttpSegment // сегменты узла
	AvgTirads Tirads
	Ai bool
}

type InsertUziRequest struct {
	Uzi Uzi
	Images []HttpImage
	Formations []HttpFormation
}