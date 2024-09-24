package models

const TiradsName = "tirads"

type Tirads struct {
	Id      int `gorm:"primaryKey"`
	Tirads1 float64
	Tirads2 float64
	Tirads3 float64
	Tirads4 float64
	Tirads5 float64
}

func (Tirads) TableName() string {
	return TiradsName
}
