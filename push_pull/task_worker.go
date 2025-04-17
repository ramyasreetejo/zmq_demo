// Task Worker
// Connects PULL socket to tcp://localhost:5557
// Collects workloads from ventilator via that socket
// Connects PUSH socket to tcp://localhost:5558
// Sends results to sink via that socket
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	// Socket to receive messages (tasks) from ventilator
	receiver, err := zmq4.NewSocket(zmq4.PULL)
	if err != nil {
		log.Fatal(err)
	}
	defer receiver.Close()
	receiver.Connect("tcp://localhost:5557")

	// Socket to send results to task sink
	sender, err := zmq4.NewSocket(zmq4.PUSH)
	if err != nil {
		log.Fatal(err)
	}
	defer sender.Close()
	sender.Connect("tcp://localhost:5558")

	workerPID := os.Getpid()

	// Process tasks continuously
	for {
		// Receive workload
		msg, err := receiver.Recv(0)
		if err != nil {
			log.Fatal(err)
		}

		// Simulate processing workload
		msec, err := strconv.ParseInt(msg, 10, 64)
		if err != nil {
			log.Printf("Invalid workload: %s\n", msg)
			continue
		}

		fmt.Printf("Worker [%d] processing task: %s ms\n", workerPID, msg)

		time.Sleep(time.Duration(msec) * time.Millisecond)

		// Send empty result message to the sink
		_, err = sender.Send("", 0)
		if err != nil {
			log.Fatal(err)
		}
	}
}
