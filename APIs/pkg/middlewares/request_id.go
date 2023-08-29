package middlewares

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", uuid.New().String())
		next.ServeHTTP(w, r)
	})
}
