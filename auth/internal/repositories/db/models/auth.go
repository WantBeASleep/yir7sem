package models

type AuthInfo struct {
	ID               uint64 `gorm:"primaryKey"`
	Login            string `gorm:"unique"`
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerID      uint64 `gorm:"unique"`
}
