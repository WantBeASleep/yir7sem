package uzi

import "github.com/google/uuid"

type Device struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Tirads struct {
	Tirads23 float64 `json:"tirads23"`
	Tirads4  float64 `json:"tirads4"`
	Tirads5  float64 `json:"tirads5"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Image struct {
	Id   uuid.UUID `json:"id"`
	Url  string    `json:"url"`
	Page int       `json:"page"`
}

type Formation struct {
	Id uuid.UUID `json:"id"`

	Ai     bool   `json:"ai"`
	Tirads Tirads `json:"tirads"`
}

type FormationReq struct {
	Ai     bool   `json:"ai"`
	Tirads Tirads `json:"tirads"`
}

type FormationNestedSegmentReq struct {
	Segments []SegmentNestedToFormation `json:"segments"`
	Ai       bool                       `json:"ai"`
	Tirads   Tirads                     `json:"tirads"`
}

type Segment struct {
	Id          uuid.UUID `json:"id"`
	ImageID     uuid.UUID `json:"image_id"`
	FormationID uuid.UUID `json:"formation_id"`

	Contor []Point `json:"contor"`
	Tirads Tirads  `json:"tirads"`
}

type SegmentUpdate struct {
	Tirads Tirads `json:"tirads"`
}

type SegmentNestedToFormation struct {
	ImageID uuid.UUID `json:"image_id"`

	Contor []Point `json:"contor"`
	Tirads Tirads  `json:"tirads"`
}

type Uzi struct {
	Id         uuid.UUID `json:"id"`
	Url        string    `json:"url"`
	Projection string    `json:"projection"`
	PatientID  uuid.UUID `json:"patient_id"`
	DeviceID   int       `json:"device_id"`
}

type UziReq struct {
	Projection string    `json:"projection"`
	PatientID  uuid.UUID `json:"patient_id"`
	DeviceID   int       `json:"device_id"`
}

type Report struct {
	Uzi        Uzi         `json:"uzi"`
	Images     []Image     `json:"images"`
	Formations []Formation `json:"formations"`
	Segments   []Segment   `json:"segments"`
}

// специфичное

type FormationWithSegments struct {
	Formation Formation `json:"formation"`
	Segments  []Segment `json:"segments"`
}

type FormationWithSegmentsIDs struct {
	FormationID uuid.UUID   `json:"formation_id"`
	SegmentsIDs []uuid.UUID `json:"segments_ids"`
}

type ImageWithSegmentsFormations struct {
	Image      Image       `json:"image"`
	Formations []Formation `json:"formations"`
	Segments   []Segment   `json:"segments"`
}
