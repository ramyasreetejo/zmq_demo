package main

import (
	"fmt"
	"log"

	"github.com/pebbe/zmq4"
)

func main() {
	// Create a REQ (Request) socket
	requester, err := zmq4.NewSocket(zmq4.REQ)
	if err != nil {
		log.Fatal(err)
	}
	defer requester.Close()

	// Connect to the server
	err = requester.Connect("tcp://localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	// Send a request
	msg := "Hello Client 2"
	fmt.Println("Sending:", msg)
	_, err = requester.Send(msg, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Receive a reply
	reply, err := requester.Recv(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received:", reply)
}
