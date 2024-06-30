package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:           "127.0.0.1:8000",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
