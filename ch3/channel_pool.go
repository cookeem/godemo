// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"
import (
	"time"
	"runtime"
)

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	numOfWorkers := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfWorkers)
	for w := 1; w <= numOfWorkers; w++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for j := range jobs {
				fmt.Println("worker", id, "started  job", j, time.Now())
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finished job", j, time.Now())
				results <- j * 2
			}
		}(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
