package models

const DeviceName = "devices"

type Device struct {
	Id   int `gorm:"primaryKey"`
	Name string
}

func (Device) TableName() string {
	return DeviceName
}
