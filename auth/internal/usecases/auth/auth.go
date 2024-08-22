package usecases

import (
	"yir/auth/internal/usecases/repositories"
)

type AuthUseCase struct {
	authRepo repositories.Auth
}

func NewAuthUseCase(
	authRepo repositories.Auth,
) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
	}
}
