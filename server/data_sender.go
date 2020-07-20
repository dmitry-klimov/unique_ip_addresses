package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var processedData <-chan uint64 = nil

func handleDataSendRequest(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "GET" {
		var lastValue uint64 = 0
		for {
			select {
			case lastValue = <-processedData:
			case <-time.After(time.Millisecond * 5):
				if _, err := w.Write([]byte(fmt.Sprintf("%d", lastValue))); err != nil {
					log.Printf("Error while sending answer, err: %s", err)
				}
			}
		}
	}
}

func StartSender(address string, processed_data <-chan uint64) *http.ServeMux {
	processedData = processed_data
	srv := http.NewServeMux()
	srv.HandleFunc("/", handleDataSendRequest) // each request calls handler
	_, _ = fmt.Printf("Listening http://%s\n", address)
	return srv
}
