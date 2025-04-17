package main

import (
	"fmt"
	"github.com/pebbe/zmq4"
	"time"
)

func main() {

	client, _ := zmq4.NewSocket(zmq4.DEALER)
	defer client.Close()
	client.SetIdentity("Client2")
	client.Connect("tcp://localhost:5555")

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Hello from Client2 - %d", i)
		client.SendMessage(msg)

		reply, _ := client.RecvMessage(0)
		fmt.Println("Client2 got:", reply)
		time.Sleep(1 * time.Second)
	}
}
