package main

import (
	"fmt"
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	// Create a REP (Reply) socket
	responder, err := zmq4.NewSocket(zmq4.REP)
	if err != nil {
		log.Fatal(err)
	}
	defer responder.Close()

	// Bind to a port
	err = responder.Bind("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server started, waiting for requests...")

	for {
		// Receive a message
		msg, err := responder.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Received:", msg)

		// Send a reply
		reply := "World"
		_, err = responder.Send(reply, 0)
		if err != nil {
			log.Fatal(err)
		}
	}
}
