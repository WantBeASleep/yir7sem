package models

type MedWorkerInfo struct {
	ID              uint64 `gorm:"primaryKey"`
	FirstName       string
	LastName        string
	MiddleName      string
	MedOrganization string
	Job             string
	IsRemoteWorker  bool `gorm:"not null;default:false"`
	ExpertDetails   string
}
