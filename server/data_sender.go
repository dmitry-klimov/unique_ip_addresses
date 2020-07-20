package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var processedData <-chan string = nil

func handleDataSendRequest(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) == "POST" {
		if err := r.ParseForm(); err == nil {
			fmt.Println("Request is processed")
		} else {
			msg := fmt.Sprintf("Error during request handling: %s\n", err.Error())
			log.Println(msg)
			if i, e := fmt.Fprintf(w, msg); e != nil {
				log.Printf("Error while sending answer, %d bytes sent, err: %s", i, e)
			}
			return
		}
	} else {
		for {
			select {
			case data := <-processedData:
				if _, err := w.Write([]byte(data)); err != nil {
					log.Printf("Error while sending answer, err: %s", err)
				}
			case <-time.After(time.Millisecond * 5):
				return
			}
		}
	}
}

func StartSender(address string, processed_data <-chan string) *http.ServeMux {
	processedData = processed_data
	srv := http.NewServeMux()
	srv.HandleFunc("/", handleDataSendRequest) // each request calls handler
	_, _ = fmt.Printf("Listening http://%s\n", address)
	return srv
}
