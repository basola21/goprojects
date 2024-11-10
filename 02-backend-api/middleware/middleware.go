package middleware

import (
	"net/http"
	"time"
)

func SimpleRateLimiter(next http.Handler, interval time.Duration) http.Handler {
	ticker := time.Tick(interval)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-ticker             // wait for the next tick before allowing the request
		next.ServeHTTP(w, r) // pass the request to the next handler
	})
}
