package register

import (
	"context"
	"fmt"

	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/repository/entity"
	"auth/internal/services/password"

	"github.com/google/uuid"
)

type Service interface {
	Register(ctx context.Context, user domain.User) (uuid.UUID, error)
}

type service struct {
	dao         repository.DAO
	passwordSrv password.Service
}

func New(
	dao repository.DAO,
	passwordSrv password.Service,
) Service {
	return &service{
		dao:         dao,
		passwordSrv: passwordSrv,
	}
}

func (s *service) Register(ctx context.Context, user domain.User) (uuid.UUID, error) {
	user.Id = uuid.New()

	salt := s.passwordSrv.GenerateSalt(ctx)
	pass, err := s.passwordSrv.Hash(ctx, user.Password, salt)
	if err != nil {
		return uuid.Nil, fmt.Errorf("hashing pass: %w", err)
	}
	user.Password = pass

	if err := s.dao.NewUserQuery(ctx).InsertUser(entity.User{}.FromDomain(user)); err != nil {
		return uuid.Nil, fmt.Errorf("insert user: %w", err)
	}

	return user.Id, nil
}
