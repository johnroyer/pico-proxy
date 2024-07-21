package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SensorData struct {
	temperature float64
	humidity    float64
}

type listen struct {
	address string
	port    int32
}

type forward struct {
	url        string
	httpMethod string
}

func main() {
	var listenOn listen

	configData, err := ini.Load("config.ini")
	if err != nil {
		// failed to load config file, use default value
		listenOn.address = "127.0.01:8000"
		listenOn.port = 8080
	} else {
		listenSection := configData.Section("listen")
		listenOn.address = listenSection.Key("address").String()

		configPort, convertErr := listenSection.Key("port").Int()
		if convertErr != nil {
			fmt.Println("port in config file is not valid")
		}
		if 1 > configPort || 65535 < configPort {
			fmt.Println("port in config file is not valid")
		}

		listenOn.port = int32(configPort)
	}

	server := http.Server{
		Addr:         listenOn.address + ":" + strconv.Itoa(int(listenOn.port)),
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
