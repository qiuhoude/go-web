package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	connectNats()

}

func connectNats() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected suc")
	defer nc.Close()

}
