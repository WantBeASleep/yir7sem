package entity

import (
	"github.com/google/uuid"
)

type TokensPair struct {
	AccessToken  string
	RefreshToken string
}

type RequestRegister struct {
	Mail        string
	LastName    string
	FirstName   string
	FathersName string
	MedOrg      string
	Password    string
}

type User struct {
	Id               uuid.UUID
	Mail             string
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerID      uuid.UUID
}
