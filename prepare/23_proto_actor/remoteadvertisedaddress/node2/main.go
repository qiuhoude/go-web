package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/examples/remotebenchmark/messages"
	"github.com/AsynkronIT/protoactor-go/remote"
)

func main() {
	remote.Start("127.0.0.1:8080")
	rootContext := actor.EmptyRootContext
	props := actor.
		PropsFromFunc(
			func(context actor.Context) {
				switch context.Message().(type) {
				case *messages.Ping:
					fmt.Printf("Received ping from sender with address:%v \n", context.Sender())
					context.Respond(&messages.Pong{})
				}
			})

	rootContext.SpawnNamed(props, "remote")

	console.ReadLine()
}
