package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"todo-api/internal/model"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const userContextKey = contextKey("username")

var JWTSecretKey = []byte(os.Getenv("JWT_SECRET"))

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			res := model.ApiResponse{}
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				w.WriteHeader(http.StatusUnauthorized)
				res.Message = "authorization header required"
				json.NewEncoder(w).Encode(res)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				res.Message = "invalid Authorization header format"
				json.NewEncoder(w).Encode(res)
				return
			}

			tokenStr := parts[1]
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
				// 서명 방법 검증
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return JWTSecretKey, nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				res.Message = "nvalid token"
				json.NewEncoder(w).Encode(res)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				res.Message = "invalid token claims"
				json.NewEncoder(w).Encode(res)
				return
			}

			// username 클레임 추출
			username, ok := claims["username"].(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				res.Message = "invalid username in token"
				json.NewEncoder(w).Encode(res)
				return
			}

			// Context에 username 저장
			ctx := context.WithValue(r.Context(), userContextKey, username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Context에서 username 가져오기
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userContextKey).(string)
	return id, ok
}
