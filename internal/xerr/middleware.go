package xerr

import (
	"encoding/json"
	"net/http"

	"user-api/internal/types"
)

func GlobalErrorMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Nếu panic là APIError
				if apiErr, ok := rec.(*APIError); ok {
					writeError(w, apiErr.Code, apiErr.Message)
					return
				}

				// default panic
				writeError(w, 500, "Lỗi không xác định")
			}
		}()

		// Gọi handler
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
