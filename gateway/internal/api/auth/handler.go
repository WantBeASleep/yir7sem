package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	adapters "gateway/internal/adapters"

	authpb "gateway/internal/generated/grpc/client/auth"
	medpb "gateway/internal/generated/grpc/client/med"
)

type Handler struct {
	adapter adapters.Adapter
}

func New(
	adapter adapters.Adapter,
) *Handler {
	return &Handler{
		adapter: adapter,
	}
}

// Register Зарегистрирует пользователя (врача)
//
//	@Summary		Зарегистрирует пользователя (врача)
//	@Description	Зарегистрирует пользователя (врача)
//	@Tags			auth
//	@Produce		json
//	@Param			body	body		RegisterIn	true	"регистрационные данные"
//	@Success		200		{object}	RegisterOut	"user_id"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	// нужна распред транзакция
	authResp, err := h.adapter.AuthAdapter.Register(ctx, &authpb.RegisterIn{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	_, err = h.adapter.MedAdapter.RegisterDoctor(ctx, &medpb.RegisterDoctorIn{
		Doctor: &medpb.Doctor{
			Id:       authResp.Id,
			Fullname: req.Fullname,
			Org:      req.Org,
			Job:      req.Job,
			Desc:     req.Desc,
		},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(authResp.Id); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// Login Получить JWT ключи
//
//	@Summary		Получить JWT ключи
//	@Description	Получить JWT ключи
//	@Tags			auth
//	@Produce		json
//	@Param			body	body		LoginIn		true	"почта, пароль"
//	@Success		200		{object}	LoginOut	"key pairs"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req LoginIn
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	authResp, err := h.adapter.AuthAdapter.Login(ctx, &authpb.LoginIn{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(authResp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}

// Refresh Получить JWT ключи
//
//	@Summary		Получить JWT ключи
//	@Description	Получить JWT ключи
//	@Tags			auth
//	@Produce		json
//	@Param			token	header		string		true	"refresh token"
//	@Success		200		{object}	RefreshOut	"key pairs"
//	@Failure		500		{string}	string		"Internal Server Error"
//	@Router			/auth/refresh [post]
func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	refreshKey := r.Header.Get("token")

	authResp, err := h.adapter.AuthAdapter.Refresh(ctx, &authpb.RefreshIn{
		RefreshToken: refreshKey,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(authResp); err != nil {
		http.Error(w, fmt.Sprintf("что то пошло не так: %v", err), 500)
		return
	}
}
