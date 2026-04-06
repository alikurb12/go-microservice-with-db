package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr: ":8080",
		Handler: buildHTTPHandler(),
	}
	log.Printf("Running Server on %s", srv.Addr)
	srv.ListenAndServe()
}