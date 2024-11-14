package uzidto

import (
	"yir/gateway/internal/entity/uzimodel"
	entity "yir/gateway/internal/entity/uzimodel"

	"github.com/google/uuid"
)

// здесь нейминг структур будет совпадать с swagger

type Formation struct {
	Id uuid.UUID `json:"id" validate:"required"`

	// meta
	Ai     bool             `json:"ai" validate:"required"`
	Tirads *uzimodel.Tirads `json:"tirads,omitempty" validate:"omitempty"`
}

type Segment struct {
	Id          uuid.UUID `json:"id" validate:"required"`
	ImageID     uuid.UUID `json:"image_id" validate:"required"`
	FormationID uuid.UUID `json:"formation_id" validate:"required"`

	// meta
	ContorURL string           `json:"contor_url" validate:"required,url"`
	Tirads    *uzimodel.Tirads `json:"tirads,omitempty" validate:"omitempty"`
}

type Report struct {
	Uzi        *uzimodel.Uzi  `json:"uzi" validate:"required"`
	Images     []entity.Image `json:"images" validate:"required,dive"`
	Formations []Formation    `json:"formations" validate:"required,dive"`
	Segments   []Segment      `json:"segments" validate:"required,dive"`
}

type FormationWithSegments struct {
	Formation *Formation `json:"formation" validate:"required"`
	Segments  []Segment  `json:"segments" validate:"required,dive"`
}

type FormationWithSegmentsIDs struct {
	FormationID uuid.UUID  `json:"formation_id"`
	SegmentsIDs uuid.UUIDs `json:"segments_ids"`
}

type ImageWithSegmentsFormations struct {
	Image      *entity.Image `json:"image" validate:"required"`
	Formations []Formation   `json:"formations" validate:"required,dive"`
	Segments   []Segment     `json:"segments" validate:"required,dive"`
}
