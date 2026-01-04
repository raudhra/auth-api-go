package middleware

import (
	"go-auth-api/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Header is empty: Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.Replace("Bearer ", "", 1)

		token, err := utils.VerifyJWT(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Unautorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
