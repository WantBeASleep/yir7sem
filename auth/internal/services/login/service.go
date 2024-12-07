package login

import (
	"context"
	"errors"
	"fmt"

	"auth/internal/repository"
	"auth/internal/repository/entity"
	"auth/internal/services/password"
	"auth/internal/services/tokenizer"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, string, error)
}

type service struct {
	dao           repository.DAO
	passwordSrv   password.Service
	tokenaizerSrv tokenizer.Service
}

func New(
	dao repository.DAO,
	passwordSrv password.Service,
	tokenaizerSrv tokenizer.Service,
) Service {
	return &service{
		dao:           dao,
		passwordSrv:   passwordSrv,
		tokenaizerSrv: tokenaizerSrv,
	}
}

func (s *service) Login(ctx context.Context, email, password string) (string, string, error) {
	userQuery := s.dao.NewUserQuery(ctx)
	userDB, err := userQuery.GetUserByEmail(email)
	if err != nil {
		return "", "", fmt.Errorf("get user by email: %w", err)
	}
	user := userDB.ToDomain()

	salt, err := s.passwordSrv.GetSalt(ctx, user.Password)
	if err != nil {
		return "", "", fmt.Errorf("get salt: %w", err)
	}

	hash, err := s.passwordSrv.Hash(ctx, password, salt)
	if err != nil {
		return "", "", fmt.Errorf("hash pass: %w", err)
	}

	if hash != user.Password {
		return "", "", errors.New("hash not equal") // TODO: починить 500тки, возвращать норм ошибки
	}

	access, refresh, err := s.tokenaizerSrv.GeneratePair(
		ctx,
		map[string]any{"x-user_id": user.Id.String()},
	)
	if err != nil {
		return "", "", fmt.Errorf("get tokens: %w", err)
	}

	user.Token = &refresh
	if _, err := userQuery.UpdateUser(entity.User{}.FromDomain(user)); err != nil {
		return "", "", fmt.Errorf("update user: %w", err)
	}

	return access, refresh, nil
}
