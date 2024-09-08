package models

const ImageName = "images"

type Image struct {
	Uuid string `gorm:"primaryKey"`
	Url  string
	Page int64

	UziUUID string
	Uzis    Uzi `gorm:"foreignKey:UziUUID;references:Uuid;"`
}

func (Image) TableName() string {
	return ImageName
}
