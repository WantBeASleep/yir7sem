package models

import "github.com/google/uuid"

const FormationName = "formations"

type Formation struct {
	Id uuid.UUID `gorm:"primaryKey"`
	Ai bool

	TiradsId int
	Tirads   Tirads `gorm:"foreignKey:TiradsId;references:Id;" copier:"-"`
}

func (Formation) TableName() string {
	return FormationName
}
