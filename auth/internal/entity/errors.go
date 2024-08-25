package entity

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPassHashNotEqual = errors.New("password hash not equal")
	ErrTokenExpired     = errors.New("jwt token expired")
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredSession   = errors.New("expired session")
)
