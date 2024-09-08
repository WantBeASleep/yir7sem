package models

const DeviceName = "devices"

type Device struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

func (Device) TableName() string {
	return DeviceName
}
