package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pebbe/zmq4"
)

func main() {
	// Create a new ZeroMQ subscriber socket
	subscriber, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		log.Fatal(err)
	}
	defer subscriber.Close()

	// Default filter (ZIP code)
	filter := "59937"

	// Get ZIP code from command line argument
	if len(os.Args) > 1 { // Usage: ./wuclient 85678
		filter = os.Args[1]
	}

	fmt.Printf("Collecting updates from weather server for %s…\n", filter)

	// Subscribe to messages that start with the given ZIP code
	err = subscriber.SetSubscribe(filter)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the weather update publisher
	err = subscriber.Connect("tcp://localhost:5556")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")

	totalTemp := 0

	// Receive 10 updates
	for i := 0; i < 10; i++ {
		// Receive a message
		message, err := subscriber.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Received msg:", message)

		// Split message into parts
		parts := strings.Split(message, " ")
		if len(parts) < 2 {
			continue // Skip malformed messages
		}

		// Parse temperature
		temp, err := strconv.Atoi(parts[1])
		if err == nil {
			totalTemp += temp
		}
	}

	// Compute and display average temperature
	fmt.Printf("Average temperature for zipcode %s was %d°F\n", filter, totalTemp/100)
}
