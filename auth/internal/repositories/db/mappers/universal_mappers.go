package mappers

import (
	"service/auth/internal/entity"
	"service/auth/internal/repositories/db/models"
)

func UserToAuthInfo(user *entity.User) (*models.AuthInfo, error) {
	return &models.AuthInfo{
		ID:               uint64(user.ID),
		Login:            user.Login,
		PasswordHash:     user.PasswordHash,
		RefreshTokenWord: user.RefreshTokenWord,
		MedWorkerID:      uint64(user.MedWorkerID),
	}, nil
}

func AuthInfoToUser(auth *models.AuthInfo) (*entity.User, error) {
	return &entity.User{
		// переполнение?
		ID:               int(auth.ID),
		Login:            auth.Login,
		PasswordHash:     auth.PasswordHash,
		RefreshTokenWord: auth.RefreshTokenWord,
		MedWorkerID:      int(auth.MedWorkerID),
	}, nil
}
