package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/alikurb12/go-microservice-with-db/internal/domain"
)

type datetimeResponce struct {
	DateTime string `json:"datatime"`
}

func CreateDateTimeHandler(dtService *domain.DateTimeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := RequestIDFromContext(r.Context())
		log.Printf(
			"new dt request method=%s path=%s requestid=%s",
			r.Method,
			r.URL.Path,
			requestID,
		)

		ts := dtService.CurrentUnixTimestamp()

		t := time.Unix(ts, 0).UTC()
		iso8601 := t.Format(time.RFC3339)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(datetimeResponce{
			DateTime: iso8601,
		})
	}
}
