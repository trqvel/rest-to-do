package api

import (
	"net/http"
	"strings"

	"github.com/trqvel/rest-to-do/source/auth"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Требуется авторизация", http.StatusUnauthorized)
			return
		}
		token := strings.Split(authHeader, " ")[1]
		_, err := auth.ValidateJWT(token)
		if err != nil {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
