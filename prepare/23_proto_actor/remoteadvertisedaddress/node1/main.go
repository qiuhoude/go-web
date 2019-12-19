package main

import (
	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/examples/remotebenchmark/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
	"log"
)

func main() {
	remote.Start("127.0.0.1:8081", remote.WithAdvertisedAddress("localhost:8081"))
	remotePid := actor.NewPID("127.0.0.1:8080", "remote")

	rootContext := actor.EmptyRootContext
	props := actor.
		PropsFromFunc(func(context actor.Context) {
			switch context.Message().(type) {
			case *actor.Started:
				message := &messages.Ping{}
				context.Request(remotePid, message)

			case *messages.Pong:
				log.Println("Received pong from sender")

			}
		})

	rootContext.Spawn(props)

	console.ReadLine()
}
