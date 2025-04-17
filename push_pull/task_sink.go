// Task Sink
// Binds PULL socket to tcp://localhost:5558
// Collects results from workers via that socket
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	//// Create a ZeroMQ context
	//context, err := zmq4.NewContext()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer context.Term()

	// Socket to receive messages from workers
	receiver, err := zmq4.NewSocket(zmq4.PULL)
	if err != nil {
		log.Fatal(err)
	}
	defer receiver.Close()
	receiver.Bind("tcp://*:5558")

	// Wait for the start-of-batch message
	msg, err := receiver.Recv(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received Start Msg:", msg)

	// Start timer
	startTime := time.Now()

	// Process 20 confirmations
	for i := 0; i < 20; i++ {
		_, err = receiver.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sink received task completion %d/20\n", i+1)
	}

	// Calculate and report elapsed time
	elapsed := time.Since(startTime).Milliseconds()
	fmt.Printf("\nTotal elapsed time: %d msec\n", elapsed)
}
