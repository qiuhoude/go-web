package main

import (
	"fmt"

	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/actor/middleware"
)

type myActor struct {
	NameAwareHolder
}

func (state *myActor) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *actor.Started:
		// this actor have been initialized by the receive pipeline
		fmt.Printf("My name is %v\n", state.name)
	}
}

type NameAware interface {
	SetName(name string)
}

type NameAwareHolder struct {
	name string
}

func (state *NameAwareHolder) SetName(name string) {
	state.name = name
}

type NamerPlugin struct{}

func (p *NamerPlugin) OnStart(ctx actor.ReceiverContext) {
	if p, ok := ctx.Actor().(NameAware); ok {
		p.SetName("GAM")
	}
}
func (p *NamerPlugin) OnOtherMessage(ctx actor.ReceiverContext, env *actor.MessageEnvelope) {}

func main() {
	rootContext := actor.EmptyRootContext
	props := actor.
		PropsFromProducer(func() actor.Actor { return &myActor{} }).
		WithReceiverMiddleware(
			middleware.Logger,
			//plugin.Use(&NamerPlugin{}),
		)

	pid := rootContext.Spawn(props)
	rootContext.Send(pid, "bar")
	console.ReadLine()
}
