package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	PubKey *rsa.PublicKey
}

func (am *AuthMiddleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получите токен из заголовка или других частей запроса
		authString := r.Header.Get("Authorization")
		tokenString := strings.Split(authString, " ")[1] // Избавляемся от Bearer
		if tokenString == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		// Разбираем и валидируем токен с помощью публичного ключа из структуры
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return am.PubKey, nil
		})

		// Проверяем результат парсинга и валидности токена
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Если токен действителен, продолжаем выполнение
		next.ServeHTTP(w, r)
	})
}
