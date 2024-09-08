package models

const ImagesFormationName = "images_formations"

type ImageFormation struct {
	Id        int64 `gorm:"primaryKey"`
	ContorURL string

	FormationUUID string
	Formations    Formation `gorm:"foreignKey:FormationUUID;references:Uuid;"`

	ImageUUID string
	Images    Image `gorm:"foreignKey:ImageUUID;references:Uuid;"`

	TiradsId string
	Tirads   Tirads `gorm:"foreignKey:TiradsId;references:Id;"`
}

func (ImageFormation) TableName() string {
	return ImagesFormationName
}
