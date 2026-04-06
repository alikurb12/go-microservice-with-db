package main

import (
	"net/http"

	"github.com/alikurb12/go-microservice-with-db/internal/domain"
	"github.com/alikurb12/go-microservice-with-db/internal/infrastructure"
	"github.com/alikurb12/go-microservice-with-db/internal/protocol/httpapi"
)

func buildHTTPHandler() http.Handler {
	idGen := infrastructure.UUIDGenerator{}
	clock := infrastructure.SystemClock{}
	systemClockService := domain.NewDateTimeService(clock)

	mux := http.NewServeMux()
	mux.Handle("GET /", httpapi.CreateDateTimeHandler(systemClockService))
	
	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)

	return handler
}