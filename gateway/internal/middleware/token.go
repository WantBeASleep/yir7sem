package middleware

import (
	"net/http"
	"yir/gateway/internal/usecase/jwt"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получите токен из заголовка или других частей запроса
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}
		// TODO: Надо переделать под единоразовую загрузку!
		// Временное решение
		pubkey, err := jwt.LoadPublicKey()
		if err != nil {
			panic("EAT SHIT BRO")
		}
		if err = jwt.ValidateJWT(token, pubkey); err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Если токен валиден, продолжайте обработку запроса
		w.Write([]byte("Middleware: Token is verified!\n"))
		next.ServeHTTP(w, r) // Вызов следующего обработчика
	})
}
