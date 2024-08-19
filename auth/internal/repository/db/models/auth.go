package models

type AuthInfo struct {
	ID uint `gorm:"primaryKey"`
	Login string `gorm:"unique"`
	PasswordHash string
	RefreshToken string
	MedWorkerID uint `gorm:"unique"`
}