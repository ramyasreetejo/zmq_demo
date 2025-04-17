package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {

	frontend, err := zmq4.NewSocket(zmq4.ROUTER)
	if err != nil {
		log.Fatal(err)
	}
	defer frontend.Close()
	frontend.Bind("tcp://*:5555") // Clients connect here

	backend, err := zmq4.NewSocket(zmq4.DEALER)
	if err != nil {
		log.Fatal(err)
	}
	defer backend.Close()
	backend.Bind("tcp://*:5556") // Workers connect here

	// Set up poller
	poller := zmq4.NewPoller()
	poller.Add(frontend, zmq4.POLLIN)
	poller.Add(backend, zmq4.POLLIN)

	fmt.Println("Broker started...")

	for {
		polled, err := poller.Poll(-1)
		if err != nil {
			log.Fatal(err)
		}

		for _, item := range polled {
			switch socket := item.Socket; socket {
			case frontend:
				// Message from client to be forwarded to worker
				msg, err := frontend.RecvMessage(0)
				if err != nil {
					log.Println("Error receiving from client:", err)
					continue
				}
				// Forward to backend (worker)
				fmt.Println("BROKER → WORKER:", msg)
				backend.SendMessage(msg)

			case backend:
				// Message from worker to be forwarded to client
				msg, err := backend.RecvMessage(0)
				if err != nil {
					log.Println("Error receiving from worker:", err)
					continue
				}
				// Forward to frontend (client)
				fmt.Println("BROKER → CLIENT:", msg)
				frontend.SendMessage(msg)
			}
		}

		time.Sleep(10 * time.Millisecond) // reduce CPU usage slightly
	}
}
