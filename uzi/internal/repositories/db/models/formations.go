package models

const FormationName = "formations"

type Formation struct {
	Uuid string `gorm:"primaryKey;autoIncrement:false"`
	Ai   bool

	TiradsId string
	Tirads   Tirads `gorm:"foreignKey:TiradsId;references:Id;"`
}

func (Formation) TableName() string {
	return FormationName
}
