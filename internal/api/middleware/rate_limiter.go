package middleware

import (
	"fmt"
	"net/http"

	"sei-ratelimiter/internal/algorithms"
)

func RateLimiter(limiter algorithms.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			clientID := r.RemoteAddr

			allowed, remaining, err := limiter.Allow(r.Context(), clientID)

			if err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))

			if !allowed {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
