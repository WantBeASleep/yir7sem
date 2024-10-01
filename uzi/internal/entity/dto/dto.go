package dto

import (
	"yir/uzi/internal/entity"

	"github.com/google/uuid"
)

// здесь нейминг структур будет совпадать с swagger

type Formation struct {
	Id uuid.UUID

	// meta
	Ai     bool
	Tirads *entity.Tirads
}

type Segment struct {
	Id          uuid.UUID
	ImageID     uuid.UUID
	FormationID uuid.UUID

	// meta
	ContorURL string
	Tirads    *entity.Tirads
}

type FormationWithSegments struct {
	Formation *Formation
	Segments  []Segment
}

type ImageWithSegmentsFormations struct {
	Image      *entity.Image
	Formations []Formation
	Segments   []Segment
}

type Uzi struct {
	UziInfo    *entity.Uzi
	Images     []entity.Image
	Formations []Formation
	Segments   []Segment
}
