package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// [POST] - /truphone/families - 105
		log.Printf("[%v] - %v - %v", r.Method, r.URL, r.ContentLength)
		next.ServeHTTP(w, r)
	})
}
