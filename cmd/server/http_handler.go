package main

import (
	"net/http"

	"github.com/alikurb12/go-microservice-with-db/internal/infrastructure"
	"github.com/alikurb12/go-microservice-with-db/internal/protocol/httpapi"
)

func buildHTTPHandler() http.Handler {
	idGen := infrastructure.UUIDGenerator{}
	mux := http.NewServeMux()
	
	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)

	return handler
}