package entity

import "github.com/google/uuid"

type RequestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DomainLogin struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type TokensPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RequestRegister struct {
	Email       string `json:"email" validate:"required,email"`
	LastName    string `json:"lastName" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	FathersName string `json:"fathersName" validate:"required"`
	MedOrg      string `json:"medOrganization" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type ResponseRegister struct {
	UUID uuid.UUID `json:"uuid"`
}

type UserCreditals struct {
	ID               int       `json:"id"`
	UUID             uuid.UUID `json:"uuid"`
	Login            string    `json:"login" validate:"required"`
	PasswordHash     string    `json:"password_hash"`
	RefreshTokenWord string    `json:"refresh_token_word"`
	MedWorkerUUID    uuid.UUID `json:"med_worker_uuid"`
}

type UserTokenVerify struct {
	UserUUID         uuid.UUID `json:"user_uuid"`
	RefreshTokenWord string    `json:"refresh_token_word"`
}
