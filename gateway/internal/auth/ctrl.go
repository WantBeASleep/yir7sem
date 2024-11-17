package auth

import (
	"context"
	pb "yir/gateway/rpc/auth"
)

type Ctrl struct {
	pb.UnimplementedAuthServer

	client pb.AuthClient
}

func NewCtrl(client pb.AuthClient) *Ctrl {
	return &Ctrl{client: client}
}

// Login godoc
//	@Summary		User Login
//	@Description	Authenticates a user and returns a token pair.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RequestLogin	true	"User Login Data"
//	@Success		200		{object}	TokensPair
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/auth/login [post]
func (c *Ctrl) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return c.client.Login(ctx, req)
}

// Register godoc
//	@Summary		User Registration
//	@Description	Registers a new user and returns a response with a UUID.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RequestRegister	true	"User Registration Data"
//	@Success		200		{object}	ResponseRegister
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/auth/register [post]
func (c *Ctrl) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return c.client.Register(ctx, req)
}

// TokenRefresh godoc
//	@Summary		Token Refresh
//	@Description	Refreshes an expired access token using a refresh token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		TokensPair	true	"Token Data"
//	@Success		200		{object}	TokensPair
//	@Failure		500		{string}	string	"Internal error"
//	@Router			/auth/token/refresh [post]
func (c *Ctrl) TokenRefresh(ctx context.Context, req *pb.TokenRefreshRequest) (*pb.TokenRefreshResponse, error) {
	return c.client.TokenRefresh(ctx, req)
}
