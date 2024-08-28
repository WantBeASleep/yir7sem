package models

const UserCreditalsName = "user_credentials"

type UserCreditals struct {
	ID               uint64 `gorm:"primaryKey"`
	UUID             string `gorm:"unique"`
	Login            string `gorm:"unique"`
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerUUID    string `gorm:"unique"`
}

func (UserCreditals) TableName() string {
	return UserCreditalsName
}
