package main

import (
	"io"
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

	mainHandler := func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Hello World\n")
	}

	http.HandleFunc("/", mainHandler)
	log.Fatal(server.ListenAndServe())
}
