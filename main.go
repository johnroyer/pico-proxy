package main

import (
	"net/http"
	"strconv"
)

func main() {
	ip := "127.0.0.1"
	port := 8080

	http.ListenAndServe(ip+strconv.Itoa(port), nil)
}
