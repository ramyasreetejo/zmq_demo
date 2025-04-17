package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	// Create a new ZeroMQ context (implicitly managed by zmq4)
	publisher, err := zmq4.NewSocket(zmq4.PUB)
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	// Bind to TCP and IPC endpoints
	err = publisher.Bind("tcp://*:5556")
	if err != nil {
		log.Fatal(err)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Infinite loop to send weather updates
	for {
		// Generate random weather data
		zipcode := 59937 + rand.Intn(2)
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		// Format message
		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity)
		fmt.Println("Sending msg:", msg)

		// Send message to all subscribers
		_, err := publisher.Send(msg, 0)
		if err != nil {
			log.Fatal(err)
		}

		// Sleep for a short duration before sending the next update
		time.Sleep(time.Millisecond * 500)
	}
}
