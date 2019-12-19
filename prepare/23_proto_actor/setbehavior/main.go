package main

import (
	"fmt"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct{ Who string }
type SetBehaviorActor struct {
	behavior actor.Behavior
}

func (state *SetBehaviorActor) Receive(context actor.Context) {
	state.behavior.Receive(context)
}

func (state *SetBehaviorActor) One(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
		state.behavior.BecomeStacked(state.Other)
	}
}

func (state *SetBehaviorActor) Other(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("%v, ey we are now handling messages in another behavior\n", msg.Who)
		state.behavior.UnbecomeStacked()
	}
}

func NewSetBehaviorActor() actor.Actor {
	act := &SetBehaviorActor{
		behavior: actor.NewBehavior(),
	}
	act.behavior.Become(act.One)
	return act
}

func main() {
	rootContext := actor.EmptyRootContext
	props := actor.PropsFromProducer(NewSetBehaviorActor)
	pid := rootContext.Spawn(props)
	rootContext.Send(pid, Hello{Who: "Roger1"})
	rootContext.Send(pid, Hello{Who: "Roger2"})
	rootContext.Send(pid, Hello{Who: "Roger3"})
	console.ReadLine()
}
