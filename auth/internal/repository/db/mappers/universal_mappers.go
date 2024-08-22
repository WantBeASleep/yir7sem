// посомтреть как делаются мапперы, пока что это полный кринжыч

package mappers

import (
	"yir/auth/internal/enity"
	"yir/auth/internal/repository/db/models"
)

func DomainUserToAuthInfo(user *enity.DomainUser) (*models.AuthInfo, error) {
	return &models.AuthInfo{
		ID:           uint64(user.ID),
		Login:        user.Login,
		PasswordHash: user.PasswordHash,
		RefreshToken: user.RefreshToken,
		MedWorkerID:  uint64(user.MedWorkerID),
	}, nil
}

func AuthInfoToDomainUser(auth *models.AuthInfo) (*enity.DomainUser, error) {
	return &enity.DomainUser{
		// переполнение?
		ID:           int(auth.ID),
		Login:        auth.Login,
		PasswordHash: auth.PasswordHash,
		RefreshToken: auth.RefreshToken,
		MedWorkerID:  int(auth.MedWorkerID),
	}, nil
}
