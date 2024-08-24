// посомтреть как делаются мапперы, пока что это полный кринжыч

package mappers

import (
	"yir/auth/internal/enity"
	"yir/auth/internal/repositories/db/models"
)

func UserToAuthInfo(user *enity.User) (*models.AuthInfo, error) {
	return &models.AuthInfo{
		ID:               uint64(user.ID),
		Login:            user.Login,
		PasswordHash:     user.PasswordHash,
		RefreshTokenWord: user.RefreshTokenWord,
		MedWorkerID:      uint64(user.MedWorkerID),
	}, nil
}

func AuthInfoToUser(auth *models.AuthInfo) (*enity.User, error) {
	return &enity.User{
		// переполнение?
		ID:               int(auth.ID),
		Login:            auth.Login,
		PasswordHash:     auth.PasswordHash,
		RefreshTokenWord: auth.RefreshTokenWord,
		MedWorkerID:      int(auth.MedWorkerID),
	}, nil
}
