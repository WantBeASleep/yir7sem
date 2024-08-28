package entity

import (
	"github.com/google/uuid"
)

// делал из расчета: при реквесте и юзкейсках содержимое меняется?
// да (pass -> passhash), 2 структуры
// нет (token pair ac rt), 1 структура

type RequestLogin struct {
	Email    string
	Password string
}

type DomainLogin struct {
	Email        string
	PasswordHash string
}

type TokensPair struct {
	AccessToken  string
	RefreshToken string
}

type RequestRegister struct {
	Email       string
	LastName    string
	FirstName   string
	FathersName string
	MedOrg      string
	Password    string
}

type ResponseRegister struct {
	UUID uuid.UUID
}

type UserCreditals struct {
	ID               int
	UUID             uuid.UUID
	Login            string
	PasswordHash     string
	RefreshTokenWord string
	MedWorkerUUID    uuid.UUID
}

type UserTokenVerify struct {
	UserUUID         uuid.UUID
	RefreshTokenWord string
}
