package middleware

import (
	"net/http"
	"time"

	"healthserver/internal/limiter"
)

var limiterInstance = limiter.NewFixedWindowLimiter(3, 10*time.Second)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !limiterInstance.Allow("hari") {
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
