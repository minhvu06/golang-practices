package xerr

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"user-api/internal/helper"
	"user-api/internal/types"

	"github.com/golang-jwt/jwt/v4"
)

// GlobalErrorMiddleware bắt lỗi toàn cục
func GlobalErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				if apiErr, ok := rec.(*APIError); ok {
					writeError(w, apiErr.Code, apiErr.Message)
					return
				}

				writeError(w, 500, "Lỗi không xác định")
			}
		}()

		next(w, r)
	}
}

// JwtMiddleware kiểm tra xác thực
var skipPaths = map[string]bool{
	"/auth":   true,
	"/health": true,
}

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if skipPaths[r.URL.Path] {
			next(w, r)
			return
		}

		auth := r.Header.Get("Authorization")
		if auth == "" {
			writeError(w, 401, "Xác thực không thành công")
			return
		}

		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			writeError(w, 401, "Xác thực không thành công")
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return helper.GetSecret(), nil
		})

		if err != nil || !token.Valid {
			writeError(w, 401, "Xác thực không thành công")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			writeError(w, 401, "Xác thực không thành công")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "claims", claims)
		r = r.WithContext(ctx)

		next(w, r)
	}
}

func writeError(w http.ResponseWriter, code int, msg string) {
	resp := types.ResponseException{
		Code:    code,
		Message: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
