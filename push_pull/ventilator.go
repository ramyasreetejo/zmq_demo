// Task ventilator
// Binds PUSH socket to tcp://localhost:5557
// Sends batch of tasks to workers via that socket
package main

import (
	"fmt"
	"log"
	"math/rand"
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

	// Socket to send messages to workers
	sender, err := zmq4.NewSocket(zmq4.PUSH)
	if err != nil {
		log.Fatal(err)
	}
	defer sender.Close()
	sender.Bind("tcp://*:5557")

	// Socket to send start-of-batch message
	sink, err := zmq4.NewSocket(zmq4.PUSH)
	if err != nil {
		log.Fatal(err)
	}
	defer sink.Close()
	sink.Connect("tcp://localhost:5558")

	// Wait for user confirmation
	fmt.Print("Press Enter when the workers are ready: ")
	fmt.Scanln()

	fmt.Println("Sending tasks to workers…")

	// Notify sink that batch processing has started
	_, err = sink.Send("0", 0)
	if err != nil {
		log.Fatal(err)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	totalMsec := 0

	for i := 0; i < 20; i++ {
		workload := rand.Intn(20)
		totalMsec += workload
		msg := fmt.Sprintf("%d", workload)
		_, err = sender.Send(msg, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Total expected cost: %d msec\n", totalMsec)

	// Give ZeroMQ time to deliver messages
	time.Sleep(time.Second)
}
