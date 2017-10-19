package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"fmt"
	"github.com/AsynkronIT/goconsole"
)

type Hi struct{ Who string }
type SetBehaviorActor struct{}

func (state *SetBehaviorActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hi:
		context.SetBehavior(state.Other)
		fmt.Printf("Hi %v\n", msg.Who)
	}
}

func (state *SetBehaviorActor) Other(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hi:
		fmt.Printf("%v, ey we are now handling messages in another behavior", msg.Who)
	}
}

func NewSetBehaviorActor() actor.Actor {
	return &SetBehaviorActor{}
}

func main() {
	props := actor.FromProducer(NewSetBehaviorActor)
	pid := actor.Spawn(props)
	pid.Tell(Hi{Who: "Haijian"})
	pid.Tell(Hi{Who: "Cookeem"})
	console.ReadLine()
}
