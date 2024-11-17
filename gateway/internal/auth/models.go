package auth

import "github.com/google/uuid"

type TokensPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestRegister struct {
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	FathersName string `json:"fathersName"`
	MedOrg      string `json:"medOrganization"`
	Password    string `json:"password"`
}

type ResponseRegister struct {
	Uuid uuid.UUID `json:"id"`
}
