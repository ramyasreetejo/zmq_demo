package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	worker, err := zmq4.NewSocket(zmq4.ROUTER)
	if err != nil {
		log.Fatal(err)
	}
	defer worker.Close()

	worker.SetIdentity("Worker1") // Needed so broker can route replies
	err = worker.Connect("tcp://localhost:5556")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Worker started...")

	for {
		// Receive message from broker
		// Expected format: [clientID, "", actual message...]
		msgParts, err := worker.RecvMessage(0)
		if err != nil {
			log.Println("Error receiving:", err)
			continue
		}

		fmt.Println("Worker received:", msgParts)

		// Simulate some work
		time.Sleep(1 * time.Second)

		// Send response back to broker: [clientID, "", response...]
		worker.SendMessage(msgParts)
	}
}
