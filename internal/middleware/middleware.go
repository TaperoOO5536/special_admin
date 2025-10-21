package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/TaperoOO5536/special_admin/internal/config"
	"github.com/TaperoOO5536/special_admin/internal/service"
)

var publicPaths = map[string]bool{
    "/v1/auth/login":   true,
    "/v1/auth/refresh": true,
    "/v1/auth/logout":  true, 
}

func AuthMiddleware(next http.Handler, authService *service.AuthService) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {        
        if publicPaths[r.URL.Path] {
            next.ServeHTTP(w, r)
            return
        }

        auth, ok := r.Header["Authorization"]
        if !ok || len(auth) == 0 {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }
        authValue := auth[0]
        if !strings.HasPrefix(strings.ToLower(authValue), strings.ToLower("Bearer")) {
            http.Error(w, "Invalid Authorization format, expected 'Bearer <token>'", http.StatusUnauthorized)
            return
        }

        token := strings.TrimPrefix(authValue, "Bearer ")
        if token == "" {
            http.Error(w, "Empty token", http.StatusUnauthorized)
            return
        }
        
        _, err := authService.Jwt.ValidateToken(token, config.GetJWTSecret())
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        ctx := context.WithValue(r.Context(), "token", token)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}