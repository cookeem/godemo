package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"fmt"
	"time"
	"github.com/AsynkronIT/goconsole"
)

type Hello2 struct{ Who string }
type Hello2Actor struct{}

func (state *Hello2Actor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		fmt.Println("Started, initialize actor here")
	case *actor.Stopping:
		fmt.Println("Stopping, actor is about shut down")
	case *actor.Stopped:
		fmt.Println("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		fmt.Println("Restarting, actor is about restart")
	case Hello2:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	props := actor.FromInstance(&Hello2Actor{})
	pid := actor.Spawn(props)
	pid.Tell(Hello2{Who: "Roger"})

	//why wait?
	//Stop is a system message and is not processed through the user message mailbox
	//thus, it will be handled _before_ any user message
	//we only do this to show the correct order of events in the console
	time.Sleep(1 * time.Second)
	pid.Stop()

	console.ReadLine()
}
