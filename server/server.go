package server

import (
	"log"
	"net/http"

	"../data_processor"
)

func Start(recv_address *string, send_address *string) int {
	received_data := make(chan string)
	processed_data := make(chan uint64)
	die := make(chan bool)

	listener := StartReceiver(*recv_address, received_data)
	sender := StartSender(*send_address, processed_data)
	go func() { log.Fatal(http.ListenAndServe(*recv_address, listener)) }()
	go func() { log.Fatal(http.ListenAndServe(*send_address, sender)) }()
	go data_processor.Start(received_data, processed_data)

	<-die
	return 0
}
