package service

import (
	"context"
	"log"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity"
	"yir/gateway/internal/pb/authpb"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthService struct {
	Client authpb.AuthClient
}

func (a *AuthService) Login(ctx context.Context, RequestLogin *entity.RequestLogin) (*entity.TokensPair, error) {
	// prepare request
	LoginRequest := &authpb.LoginRequest{
		Email:    RequestLogin.Email,
		Password: RequestLogin.Password,
	}
	// call remote procedure
	LoginResponse, err := a.Client.Login(ctx, LoginRequest)
	if err != nil {
		custom.Logger.Error(
			"failed to login",
			zap.Error(err),
		)
		return nil, err
	}
	custom.Logger.Info(
		"Account entered",
		zap.String("Email", RequestLogin.Email),
	)
	// convert to entity
	tokenPair := &entity.TokensPair{
		AccessToken:  LoginResponse.GetAccessToken(),
		RefreshToken: LoginResponse.GetRefreshToken(),
	}
	return tokenPair, nil
}

func (a *AuthService) Register(ctx context.Context, RequestRegister *entity.RequestRegister) (*entity.ResponseRegister, error) {
	RegisterRequest := &authpb.RegisterRequest{
		Email:           RequestRegister.Email,
		LastName:        RequestRegister.LastName,
		FirstName:       RequestRegister.FirstName,
		FathersName:     RequestRegister.FathersName,
		MedOrganization: RequestRegister.MedOrg,
		Password:        RequestRegister.Password,
	}

	log.Println(
		"register service:",
		RegisterRequest,
	)

	RegisterResponse, err := a.Client.Register(ctx, RegisterRequest)
	if err != nil {
		custom.Logger.Error(
			"failed to register",
			zap.Error(err),
		)
		return nil, err
	}
	custom.Logger.Info(
		"Account created",
		zap.String("Email", RequestRegister.Email),
	)
	resp := &entity.ResponseRegister{}
	resp.UUID, err = uuid.Parse(RegisterResponse.Uuid)
	if err != nil {
		custom.Logger.Error(
			"failed to parse register uuid",
			zap.Error(err),
		)
		return nil, err
	}
	return resp, nil
}

func (a *AuthService) TokenRefresh(ctx context.Context, TokensPair *entity.TokensPair) (*entity.TokensPair, error) {
	TokensRefreshReq := &authpb.TokenRefreshRequest{
		RefreshToken: TokensPair.RefreshToken,
	}
	TokenRefreshResp, err := a.Client.TokenRefresh(ctx, TokensRefreshReq)
	if err != nil {
		custom.Logger.Error(
			"failed to refresh token",
			zap.Error(err),
		)
		return nil, err
	}
	custom.Logger.Debug(
		"Token refreshed",
		zap.String("Old refresh token", TokensPair.AccessToken),
	)
	tokenRefresh := &entity.TokensPair{
		AccessToken:  TokenRefreshResp.GetAccessToken(),
		RefreshToken: TokenRefreshResp.GetRefreshToken(),
	}
	return tokenRefresh, nil
}
