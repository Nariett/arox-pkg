package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Nariett/arox-pkg/response"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Unauthorized(w, "Authorization header is required")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.BadRequest(w, "Token must use Bearer scheme")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {

			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					response.Unauthorized(w, "Token has expired")
					return
				}
			}
			response.Unauthorized(w, "Invalid token: "+err.Error())
			return
		}

		if !token.Valid {
			response.Unauthorized(w, "Invalid token")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			if exp, exists := claims["exp"]; exists {
				expTime, ok := exp.(float64)
				if !ok {
					response.Unauthorized(w, "Invalid expiration time format")
					return
				}

				if time.Now().Unix() > int64(expTime) {
					response.Unauthorized(w, "Token has expired")
					return
				}
			}

			if _, exists := claims["uuid"]; !exists {
				response.Unauthorized(w, "Token missing required claims")
				return
			}

			ctx := r.Context()

			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
