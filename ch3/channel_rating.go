package main

import (
	"time"
	"fmt"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//每500ms触发一次写入的channel
	limiter := time.Tick(time.Millisecond * 500)

	for req := range requests {
		//限速
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//#########################
	//创建一个buffer为3的channel，代替无buffer的time.Tick
	burstyLimiter := make(chan time.Time, 3)

	//预先把buffer写满
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	go func() {
		for i := 1; i <= 10; i++ {
			burstyRequests <- i
		}
		close(burstyRequests)
	}()

	for req := range burstyRequests {
		//前3个request会立刻执行，因为burstyLimiter已经有数据
		//后边的request会间隔500ms执行一个
		<-burstyLimiter
		fmt.Println("burstyRequests", req, time.Now())
	}
}
