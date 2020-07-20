package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var receivedData chan<- string = nil

func handleDataReceiveRequest(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		data := "{ \"timestamp\": \"2020-06-24T15:27:00.123456Z\", \"ip\": \"83.150.59.250\", \"url\": ... }"
		receivedData <- data
		if i, e := fmt.Fprintf(w, data); e != nil {
			log.Printf("Error while sending answer, %d bytes sent, err: %s", i, e)
		}
	} else {
		log.Printf("Wrong request")
	}
}

func StartReceiver(address string, received_data chan<- string) *http.ServeMux {
	receivedData = received_data
	srv := http.NewServeMux()
	srv.HandleFunc("/", handleDataReceiveRequest) // each request calls handler
	_, _ = fmt.Printf("Listening http://%s\n", address)
	return srv
}
