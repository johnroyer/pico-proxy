package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type SensorData struct {
	temperature float64
	humidity    float64
}

func main() {
	server := http.Server{
		Addr:           "127.0.0.1:8000",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/pico", mainHandler)

	log.Fatal(server.ListenAndServe())
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	data := SensorData{0, 0}
	if tmp, error := strconv.ParseFloat(request.URL.Query().Get("tmp"), 64); error == nil {
		data.temperature = tmp
	}
	data.temperature = tmp

	if tmp, error := strconv.ParseFloat(request.URL.Query().Get("hum"), 64); error == nil {
		data.humidity = tmp
	}
}
