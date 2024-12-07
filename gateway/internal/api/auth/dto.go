package auth

import (
	"github.com/google/uuid"
)

type JWTKeyPair struct {
	AccessKey  string `json:"access_key"`
	RefreshKey string `json:"refresh_key"`
}

type RegisterIn struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Fullname string  `json:"fullname"`
	Org      string  `json:"org"`
	Job      string  `json:"job"`
	Desc     *string `json:"desc"`
}

type RegisterOut struct {
	Id uuid.UUID `json:"id"`
}

type LoginIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOut struct {
	JWTKeyPair
}

type RefreshIn struct{}

type RefreshOut struct {
	JWTKeyPair
}
