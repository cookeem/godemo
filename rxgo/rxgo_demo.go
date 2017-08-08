package main

import (
	"errors"
	"fmt"
	"github.com/reactivex/rxgo/iterable"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {
	watcher := observer.Observer{

		// Register a handler function for every next available item.
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},

		// Register a handler for any emitted error.
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
		},

		// Register a handler when a stream is completed.
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	it, _ := iterable.New([]interface{}{1, 2, 3, 4, errors.New("bang"), 5})
	source := observable.From(it)

	sub := source.Subscribe(watcher)
	<-sub
}
