package uzi

import (
	"time"

	"github.com/google/uuid"
)

type Uzi struct {
	Id         uuid.UUID `json:"id"`
	Projection string    `json:"projection"`
	Checked    bool      `json:"checked"`
	PatientID  uuid.UUID `json:"patient_id"`
	DeviceID   int       `json:"device_id"`
	CreateAt   time.Time `json:"create_at"`
}

type Echographic struct {
	Id              uuid.UUID `json:"id"`
	Contors         *string   `json:"contors"`
	LeftLobeLength  *float64  `json:"left_lobe_length"`
	LeftLobeWidth   *float64  `json:"left_lobe_width"`
	LeftLobeThick   *float64  `json:"left_lobe_thick"`
	LeftLobeVolum   *float64  `json:"left_lobe_volum"`
	RightLobeLength *float64  `json:"right_lobe_length"`
	RightLobeWidth  *float64  `json:"right_lobe_width"`
	RightLobeThick  *float64  `json:"right_lobe_thick"`
	RightLobeVolum  *float64  `json:"right_lobe_volum"`
	GlandVolum      *float64  `json:"gland_volum"`
	Isthmus         *float64  `json:"isthmus"`
	Struct          *string   `json:"struct"`
	Echogenicity    *string   `json:"echogenicity"`
	RegionalLymph   *string   `json:"regional_lymph"`
	Vascularization *string   `json:"vascularization"`
	Location        *string   `json:"location"`
	Additional      *string   `json:"additional"`
	Conclusion      *string   `json:"conclusion"`
}

type Image struct {
	Id   uuid.UUID `json:"id"`
	Page int       `json:"page"`
}

type Node struct {
	Id       uuid.UUID `json:"id"`
	Ai       bool      `json:"ai"`
	Tirads23 float64   `json:"tirads23"`
	Tirads4  float64   `json:"tirads4"`
	Tirads5  float64   `json:"tirads5"`
}

type Segment struct {
	Id       uuid.UUID `json:"id"`
	ImageID  uuid.UUID `json:"image_id"`
	NodeID   uuid.UUID `json:"node_id"`
	Contor   string    `json:"contor"`
	Tirads23 float64   `json:"tirads23"`
	Tirads4  float64   `json:"tirads4"`
	Tirads5  float64   `json:"tirads5"`
}

type Device struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PatchUziIn struct {
	Projection *string `json:"projection"`
	Checked    *bool   `json:"checked"`
}

type PatchUziOut struct {
	Uzi
}

type PatchEchographicsIn struct {
	Contors         *string  `json:"contors"`
	LeftLobeLength  *float64 `json:"left_lobe_length"`
	LeftLobeWidth   *float64 `json:"left_lobe_width"`
	LeftLobeThick   *float64 `json:"left_lobe_thick"`
	LeftLobeVolum   *float64 `json:"left_lobe_volum"`
	RightLobeLength *float64 `json:"right_lobe_length"`
	RightLobeWidth  *float64 `json:"right_lobe_width"`
	RightLobeThick  *float64 `json:"right_lobe_thick"`
	RightLobeVolum  *float64 `json:"right_lobe_volum"`
	GlandVolum      *float64 `json:"gland_volum"`
	Isthmus         *float64 `json:"isthmus"`
	Struct          *string  `json:"struct"`
	Echogenicity    *string  `json:"echogenicity"`
	RegionalLymph   *string  `json:"regional_lymph"`
	Vascularization *string  `json:"vascularization"`
	Location        *string  `json:"location"`
	Additional      *string  `json:"additional"`
	Conclusion      *string  `json:"conclusion"`
}

type PatchEchographicsOut struct {
	Echographic
}

type GetUziIn struct{}

type GetUziOut struct {
	Uzi
}

type GetPatientUziIn struct{}

type GetPatientUziOut struct {
	Uzis []Uzi `json:"uzis"`
}

type GetEchographicsIn struct{}

type GetEchographicsOut struct {
	Echographic
}

type GetUziImagesIn struct{}

type GetUziImagesOut struct {
	Images []Image `json:"images"`
}

type GetUziNodeSegmentsIn struct{}

type GetUziNodeSegmentsOut struct {
	Nodes    []Node    `json:"nodes"`
	Segments []Segment `json:"segments"`
}

type GetUziDeviceIn struct{}

type GetUziDeviceOut struct {
	Devices []Device `json:"devices"`
}

type PostNodeIn struct {
	Segments []struct {
		ImageID  uuid.UUID `json:"image_id"`
		Contor   string    `json:"contor"`
		Tirads23 float64   `json:"tirads23"`
		Tirads4  float64   `json:"tirads4"`
		Tirads5  float64   `json:"tirads5"`
	} `json:"segments"`

	UziID    uuid.UUID `json:"uzi_id"`
	Tirads23 float64   `json:"tirads23"`
	Tirads4  float64   `json:"tirads4"`
	Tirads5  float64   `json:"tirads5"`
}

type GetAllNodesIn struct{}

type GetAllNodesOut struct {
	Nodes []Node `json:"nodes"`
}

type PostNodeOut struct {
	Id uuid.UUID `json:"id"`
}

type (
	DeleteNodeIn  struct{}
	DeleteNodeOut struct{}
)

type PatchNodeIn struct {
	Tirads23 *float64 `json:"tirads23"`
	Tirads4  *float64 `json:"tirads4"`
	Tirads5  *float64 `json:"tirads5"`
}

type PatchNodeOut struct {
	Node
}

type PostSegmentIn struct {
	ImageID  uuid.UUID `json:"image_id"`
	NodeID   uuid.UUID `json:"node_id"`
	Contor   string    `json:"contor"`
	Tirads23 float64   `json:"tirads23"`
	Tirads4  float64   `json:"tirads4"`
	Tirads5  float64   `json:"tirads5"`
}

type PostSegmentOut struct {
	Id uuid.UUID `json:"id"`
}

type (
	DeleteSegmentIn  struct{}
	DeleteSegmentOut struct{}
)

type PatchSegmentIn struct {
	Tirads23 *float64 `json:"tirads23"`
	Tirads4  *float64 `json:"tirads4"`
	Tirads5  *float64 `json:"tirads5"`
}

type PatchSegmentOut struct {
	Segment
}
