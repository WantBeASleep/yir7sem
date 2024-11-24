package models

import "github.com/google/uuid"

type MedWorkerInfo struct {
	ID              uuid.UUID `gorm:"primaryKey"`
	FirstName       string
	LastName        string
	FathersName     string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool `gorm:"not null;default:false"`
	ExpertDetails   string
}

func (MedWorkerInfo) TableName() string {
	return MedWorkersName
}
