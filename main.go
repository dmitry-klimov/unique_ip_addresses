package main

import (
	"flag"
	"fmt"
	"log"

	"./common"
	"./server"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
		fmt.Printf("%s - see you! =)", common.GetInfo())
	}()

	receiverIP := flag.String("listen_to", "127.0.0.1:5000", "address:port to receive logs")
	senderIP := flag.String("send_to", "127.0.0.1:9102", "address:port to send results")
	flag.Parse()

	fmt.Println(common.GetInfo())
	log.Fatal(server.Start(receiverIP, senderIP))
}
