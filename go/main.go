package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	NatServer = "nats://127.0.0.1:4222"
	// NatServer = "nats://nats.gondor.svc.kube:4222"
	Subject = "time"
)

func main() {
	nats.NkeyOptionFromSeed()
	nc, err := nats.Connect(NatServer)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	for {
		msg := time.Now().Format(time.RFC3339)
		if err := nc.Publish(Subject, []byte(msg)); err != nil {
			fmt.Printf("Failed to publish msg. ERR: %+v\n", err)
		}
		fmt.Println("Published 1 msg")
		time.Sleep(time.Second)
	}
}
