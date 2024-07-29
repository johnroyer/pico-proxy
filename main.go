package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type SensorData struct {
	temperature float64
	humidity    float64
}

type listen struct {
	address string
	port    int
}

type forward struct {
	url        string
	httpMethod string
}

func main() {
	if _, err := os.Stat("config.ini"); os.IsNotExist(err) {
		// config file not found
		fmt.Println("file 'config.ini' not found")
		os.Exit(1)
	}

	iniConfig, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config.ini not foud")
		os.Exit(1)
	}
	listenConfig := getListenData(iniConfig)
	forwardConfig := getForwardDate(iniConfig)

	server := http.Server{
		Addr:         listenConfig.address + ":" + strconv.Itoa(int(listenConfig.port)),
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

func getListenData(iniFile *ini.File) listen {
	var listenConfig listen
	listenConfig.address = iniFile.Section("listen").Key("address").MustString("127.0.0.1")
	listenConfig.port = iniFile.Section("listen").Key("port").InInt(8080)

	return listenConfig
}

func getForwardDate(iniFile *ini.File) forward {
	var forwardConfig forward

	forwardConfig.url = iniFile.Section("forward").Key("url").MustString("https://my.domain/api/test")
	forwardConfig.httpMethod = iniFile.Section("forward").Key("method").MustString("POST")
}
