package middleware

import (
	"crypto/rsa"
	"log/slog"
	"net/http"

	authlib "github.com/WantBeASleep/med_ml_lib/auth"
	loglib "github.com/WantBeASleep/med_ml_lib/log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type middlewares struct {
	publicKey *rsa.PublicKey
}

func New(
	publicKey *rsa.PublicKey,
) *middlewares {
	return &middlewares{
		publicKey: publicKey,
	}
}

// TODO: вынести это в библиотку med_ml_lib, и поменять парс jwt в тестах auth
// распарсит токен, положит в хедер x-user_id
func (m *middlewares) Jwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("token")
		if tokenString == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return m.publicKey, nil })
		if err != nil || !token.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["x-user_id"].(string); ok {
				r.Header.Set("x-user_id", userID)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// TODO: залогирует + сделать x-request_id (ПЕРЕДЕЛАТЬ)
func (m *middlewares) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New()
		ctx := authlib.WithRequestID(r.Context(), requestID)
		ctx = loglib.WithField(ctx, "x-request_id", requestID.String())

		slog.InfoContext(ctx, "New request", slog.String("path", r.URL.Path))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
