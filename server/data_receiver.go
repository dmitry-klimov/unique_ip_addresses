package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var data_channel chan<- string = nil

func handleDataReceiveRequest(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		fmt.Println("Request is processed")
		data_channel <- "{ \"timestamp\": \"2020-06-24T15:27:00.123456Z\", \"ip\": \"83.150.59.250\", \"url\": ... }"
	} else {
		log.Printf("Wrong request")
	}
}

func StartReceiver(address string, received_data chan<- string) *http.ServeMux {
	data_channel = received_data
	srv := http.NewServeMux()
	srv.HandleFunc("/", handleDataReceiveRequest) // each request calls handler
	_, _ = fmt.Printf("Listening http://%s\n", address)
	return srv
}
