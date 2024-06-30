package main

import (
	"fmt"
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

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.RequestURI)
	})

	http.HandleFunc("/pico", mainHandler)

	log.Fatal(server.ListenAndServe())
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	data := SensorData{0, 0}
	if tmp, error := strconv.ParseFloat(request.URL.Query().Get("tmp"), 64); error == nil {
		data.temperature = tmp
	} else {
		fmt.Println("Failed to parse temperature")
	}

	if tmp, error := strconv.ParseFloat(request.URL.Query().Get("hum"), 64); error == nil {
		data.humidity = tmp
	} else {
		fmt.Println("Failed to parse humidity")
	}
}
