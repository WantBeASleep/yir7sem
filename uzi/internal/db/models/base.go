package models

import (
	"github.com/google/uuid"
)

const (
	UziName       = "uzis"
	ImageName     = "images"
	FormationName = "formations"
	SegmentsName  = "segments"
	DeviceName    = "devices"
	TiradsName    = "tirads"
)

type Uzi struct {
	Id  uuid.UUID `gorm:"primaryKey"`
	Url string

	Projection string
	PatientID  uuid.UUID
	DeviceID   int
	Devices    Device `gorm:"foreignKey:DeviceID;references:Id" copier:"-"`
}

func (Uzi) TableName() string {
	return UziName
}

type Image struct {
	Id    uuid.UUID `gorm:"primaryKey"`
	Url   string
	UziID uuid.UUID
	Uzi   Uzi `gorm:"foreignKey:UziID;references:Id;" copier:"-"`

	Page int
}

func (Image) TableName() string {
	return ImageName
}

type Formation struct {
	Id uuid.UUID `gorm:"primaryKey"`

	Ai       bool
	TiradsID int
	Tirads   Tirads `gorm:"foreignKey:TiradsID;references:Id;" copier:"-"`
}

func (Formation) TableName() string {
	return FormationName
}

type Segment struct {
	Id          uuid.UUID `gorm:"primaryKey"`
	ImageID     uuid.UUID
	Image       Image `gorm:"foreignKey:ImageID;references:Id;" copier:"-"`
	FormationID uuid.UUID
	Formation   Formation `gorm:"foreignKey:FormationID;references:Id;" copier:"-"`

	ContorURL string
	TiradsID  int
	Tirads    Tirads `gorm:"foreignKey:TiradsID;references:Id;" copier:"-"`
}

func (Segment) TableName() string {
	return SegmentsName
}

type Device struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func (Device) TableName() string {
	return DeviceName
}

type Tirads struct {
	Id       int `gorm:"primaryKey"`
	Tirads23 float64
	Tirads4  float64
	Tirads5  float64
}

func (Tirads) TableName() string {
	return TiradsName
}
