package auth

import (
	"context"
	"net/http"
	"yir/gateway/internal/custom"
	"yir/gateway/internal/entity"
	"yir/gateway/internal/middleware"

	validator "github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, RequestLogin *entity.RequestLogin) (*entity.TokensPair, error)
	Register(ctx context.Context, RequestRegister *entity.RequestRegister) (*entity.ResponseRegister, error)
	TokenRefresh(ctx context.Context, TokensPair *entity.TokensPair) (*entity.TokensPair, error)
}

type AuthController struct {
	Service    AuthService
	Middleware *middleware.AuthMiddleware
}

func New(s AuthService) *AuthController {
	return &AuthController{
		Service: s,
	}
}

// Login godoc
// @Summary      User Login
// @Description  Authenticates a user and returns a token pair.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  entity.RequestLogin  true  "User Login Data"
// @Success      200   {object}   entity.TokensPair
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to authenticate user"
// @Router       /auth/login [post]
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	// перегонка в структуру
	var LoginData entity.RequestLogin
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(r.Body).Decode(&LoginData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// валидация
	validate := validator.New()
	if err := validate.Struct(LoginData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	// отдаем в юзкейс
	ctx := context.Background()
	tokenpair, err := c.Service.Login(ctx, &LoginData)
	if err != nil {
		http.Error(w, "Failed to authenticate user. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(&tokenpair); err != nil {
		http.Error(w, "Failed to generate response. Please try again later.", http.StatusBadGateway)
		custom.Logger.Error(
			"json encoding failed",
			zap.Error(err),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Register godoc
// @Summary      User Registration
// @Description  Registers a new user and returns a response with a UUID.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  entity.RequestRegister  true  "User Registration Data"
// @Success      200   {object}   entity.ResponseRegister
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to register user"
// @Router       /auth/register [post]
func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var RegisterData entity.RequestRegister
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.NewDecoder(r.Body).Decode(&RegisterData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// валидация
	validate := validator.New()
	if err := validate.Struct(RegisterData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	// отдаем в юзкейс
	ctx := context.Background()
	data, err := c.Service.Register(ctx, &RegisterData)
	if err != nil {
		http.Error(w, "Failed to register user. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later.", http.StatusBadGateway)
		custom.Logger.Error(
			"json encoding failed",
			zap.Error(err),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// TokenRefresh godoc
// @Summary      Token Refresh
// @Description  Refreshes an expired access token using a refresh token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  entity.TokensPair  true  "Token Data"
// @Success      200   {object}   entity.TokensPair
// @Failure      400   {string}   string  "Invalid request payload"
// @Failure      502   {string}   string  "Failed to refresh token"
// @Router       /auth/token/refresh [post]
func (c *AuthController) TokenRefresh(w http.ResponseWriter, r *http.Request) {

	var TokenData entity.TokensPair
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.NewDecoder(r.Body).Decode(&TokenData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// валидация
	validate := validator.New()
	if err := validate.Struct(TokenData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		custom.Logger.Error(
			"validation failed",
			zap.Error(err),
		)
		return
	}
	// отдаем в юзкейс
	ctx := context.Background()
	data, err := c.Service.TokenRefresh(ctx, &TokenData)
	if err != nil {
		http.Error(w, "Failed to refresh token. Please try again later.", http.StatusBadGateway)
		return
	}
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		http.Error(w, "Failed to generate response. Please try again later.", http.StatusBadGateway)
		custom.Logger.Error(
			"json encoding failed",
			zap.Error(err),
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}
