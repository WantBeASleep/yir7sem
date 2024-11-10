package models

import "github.com/google/uuid"

const UserCreditalsName = "users"

type User struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	Mail             string    `gorm:"unique"`
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerID      uuid.UUID `gorm:"unique"`
}

func (User) TableName() string {
	return UserCreditalsName
}
