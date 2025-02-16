//go:build e2e

package e2e_test

import (
	pb "auth/internal/generated/grpc/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func (suite *TestSuite) TestRefresh_Success() {
	testEmail := randstr.String(5) + "@mail.ru"
	testPassword := randstr.String(5)

	// регистрация пользователя
	_, err := suite.client.Register(
		suite.ctx,
		&pb.RegisterIn{
			Email:    testEmail,
			Password: testPassword,
		},
	)
	require.NoError(suite.T(), err)

	// получаем токены
	loginResponse, err := suite.client.Login(
		suite.ctx,
		&pb.LoginIn{
			Email:    testEmail,
			Password: testPassword,
		},
	)
	require.NoError(suite.T(), err)

	// проверяем 3 раза!
	refreshToken := loginResponse.RefreshToken
	for range 3 {
		newTokens, err := suite.client.Refresh(
			suite.ctx,
			&pb.RefreshIn{RefreshToken: refreshToken},
		)
		require.NoError(suite.T(), err)

		// access token
		token, err := jwt.Parse(newTokens.AccessToken, func(token *jwt.Token) (interface{}, error) { return suite.publicKey, nil })
		require.NoError(suite.T(), err)
		require.True(suite.T(), token.Valid)

		// refresh token
		token, err = jwt.Parse(newTokens.RefreshToken, func(token *jwt.Token) (interface{}, error) { return suite.publicKey, nil })
		require.NoError(suite.T(), err)
		require.True(suite.T(), token.Valid)

		refreshToken = newTokens.RefreshToken
	}
}
