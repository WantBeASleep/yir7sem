package mappers

import (
	"yir/auth/internal/entity"
	"yir/auth/internal/repositories/db/models"

	"github.com/google/uuid"
)

func EntityUserCreditalsToModelUserCreditals(user *entity.UserCreditals) (*models.UserCreditals, error) {
	return &models.UserCreditals{
		ID:               uint64(user.ID),
		UUID:             user.UUID.String(),
		Login:            user.Login,
		PasswordHash:     user.PasswordHash,
		RefreshTokenWord: user.RefreshTokenWord,
		MedWorkerUUID:    user.MedWorkerUUID.String(),
	}, nil
}

func ModelUserCreditalsToEntityUserCreditals(auth *models.UserCreditals) (*entity.UserCreditals, error) {
	return &entity.UserCreditals{
		ID:               int(auth.ID),
		UUID:             uuid.MustParse(auth.UUID),
		Login:            auth.Login,
		PasswordHash:     auth.PasswordHash,
		RefreshTokenWord: auth.RefreshTokenWord,
		MedWorkerUUID:    uuid.MustParse(auth.MedWorkerUUID),
	}, nil
}
