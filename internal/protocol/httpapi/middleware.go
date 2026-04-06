package httpapi

import (
	"net/http"

	"github.com/alikurb12/go-microservice-with-db/internal/domain"
)

func RequestIDMiddleware(genID domain.IDGenerator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := genID.NewID()
			w.Header().Set(RequestIdHeader, requestID)
			ctx := WithRequestIdContext(r.Context(), requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
