package middleware

import (
	"crypto/rsa"
	"log/slog"
	"net/http"

	"github.com/WantBeASleep/goooool/ctxlib"

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

// залогирует + сделать x-request_id
func (m *middlewares) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New()
		// TODO: подумать над какой то глобавльной переменной x-request_id
		ctx := ctxlib.PublicSet(r.Context(), "x-request_id", requestID.String())

		slog.InfoContext(ctx, "New request", slog.String("path", r.URL.Path))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
