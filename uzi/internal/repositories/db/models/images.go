package models

import "github.com/google/uuid"

const ImageName = "images"

type Image struct {
	Id   uuid.UUID `gorm:"primaryKey"`
	Url  string
	Page int64

	UziID uuid.UUID
	Uzi   Uzi `gorm:"foreignKey:UziID;references:Id;" copier:"-"`
}

func (Image) TableName() string {
	return ImageName
}
