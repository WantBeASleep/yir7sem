package dto

import (
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

type Formation struct {
	Id uuid.UUID

	Ai     bool
	Tirads *entity.Tirads
}

type Segment struct {
	Id          uuid.UUID
	ImageID     uuid.UUID
	FormationID uuid.UUID

	Contor string
	Tirads *entity.Tirads
}

type Report struct {
	Uzi        *entity.Uzi
	Images     []entity.Image
	Formations []Formation
	Segments   []Segment
}

type FormationWithSegments struct {
	Formation *Formation
	Segments  []Segment
}

// нужно для массивов
type FormationWithSegmentsIDs struct {
	FormationID uuid.UUID
	SegmentsIDs uuid.UUIDs
}

type ImageWithFormationsSegments struct {
	Image      *entity.Image
	Formations []Formation
	Segments   []Segment
}
