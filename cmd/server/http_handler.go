package main

import (
	"log"
	"net/http"

	"github.com/alikurb12/go-microservice-with-db/internal/domain"
	"github.com/alikurb12/go-microservice-with-db/internal/infrastructure"
	"github.com/alikurb12/go-microservice-with-db/internal/infrastructure/sqlite_clock"
	"github.com/alikurb12/go-microservice-with-db/internal/protocol/httpapi"
)

func buildHTTPHandler() http.Handler {
	idGen := infrastructure.UUIDGenerator{}
	// clock := infrastructure.SystemClock{}
	// systemClockService := domain.NewDateTimeService(clock)

	db, err := sqlite_clock.Open("app.db")
	if err != nil {
		log.Fatal(err)
	}
	sqlClock := sqlite_clock.NewSQLiteClock(db)
	sqliteClockService := domain.NewDateTimeService(sqlClock)

	mux := http.NewServeMux()
	mux.Handle("GET /", httpapi.CreateDateTimeHandler(sqliteClockService))
	
	var handler http.Handler = mux
	handler = httpapi.RequestIDMiddleware(idGen)(handler)

	return handler
}