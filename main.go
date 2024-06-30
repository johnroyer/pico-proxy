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
		Addr:         "127.0.0.1:8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.RequestURI)
	})

	http.HandleFunc("/pico", sensorDataHandler)

	log.Fatal(server.ListenAndServe())
}

func sensorDataHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.RequestURI)

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

	writer.WriteHeader(200)
}
