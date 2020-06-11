package middleware

import (
	"github.com/google/uuid"
	"github.com/ricardomgoncalves/truphone_ta_go/pkg/requestid"
	"net/http"
)

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(requestid.WithRequestId(r.Context(), uuid.New().String())))
	})
}
