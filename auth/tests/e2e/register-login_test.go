//go:build e2e

package e2e_test

import (
	pb "auth/internal/generated/grpc/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thanhpk/randstr"
)

func (suite *TestSuite) TestRegisterLogin_Success() {
	testEmail := randstr.String(5) + "@mail.ru"
	testPassword := randstr.String(5)

	// регистрация пользователя
	registerResponse, err := suite.client.Register(
		suite.ctx,
		&pb.RegisterIn{
			Email:    testEmail,
			Password: testPassword,
		},
	)
	require.NoError(suite.T(), err)

	// проверяем что uuid правильный
	userID, err := uuid.Parse(registerResponse.Id)
	require.NoError(suite.T(), err)

	// проверка логина
	loginResponse, err := suite.client.Login(
		suite.ctx,
		&pb.LoginIn{
			Email:    testEmail,
			Password: testPassword,
		},
	)
	require.NoError(suite.T(), err)

	// access token
	token, err := jwt.Parse(loginResponse.AccessToken, func(token *jwt.Token) (interface{}, error) { return suite.publicKey, nil })
	require.NoError(suite.T(), err)
	require.True(suite.T(), token.Valid)

	// refresh token
	token, err = jwt.Parse(loginResponse.RefreshToken, func(token *jwt.Token) (interface{}, error) { return suite.publicKey, nil })
	require.NoError(suite.T(), err)
	require.True(suite.T(), token.Valid)

	// ID в JWT ключе должен совпадать
	require.Equal(suite.T(), userID.String(), token.Claims.(jwt.MapClaims)["x-user_id"].(string))
}

func (suite *TestSuite) TestRegisterLogin_Fail_WrongPass() {
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

	// проверка логина
	_, err = suite.client.Login(
		suite.ctx,
		&pb.LoginIn{
			Email:    testEmail,
			Password: "WRONG_PASSWORD",
		},
	)
	require.Error(suite.T(), err)
}
