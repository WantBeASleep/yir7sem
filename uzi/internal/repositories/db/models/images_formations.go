package models

import "github.com/google/uuid"

const ImagesFormationName = "images_formations"

type ImageFormation struct {
	Id        int `gorm:"primaryKey"`
	ContorURL string

	FormationID uuid.UUID
	Formation   Formation `gorm:"foreignKey:FormationID;references:Id;" copier:"-"`

	ImageID uuid.UUID
	Image   Image `gorm:"foreignKey:ImageID;references:Id;" copier:"-"`

	TiradsId int
	Tirads   Tirads `gorm:"foreignKey:TiradsId;references:Id;" copier:"-"`
}

func (ImageFormation) TableName() string {
	return ImagesFormationName
}
