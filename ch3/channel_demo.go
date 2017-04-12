package main

import (
	"fmt"
	"time"
)

func Count(ch chan int, i int) {
	ch <- 1
	fmt.Println("Counting", i)
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}
	for _, ch := range chs {
		<- ch
	}

	time.Sleep(1 * time.Second)

	//ch := make(chan int, 1)
	//for {
	//	select {
	//	case ch <- 0:
	//		fmt.Println("ch <- 0", time.Now())
	//		time.Sleep(1 * time.Second)
	//	case ch <- 1:
	//		fmt.Println("ch <- 1", time.Now())
	//		time.Sleep(2 * time.Second)
	//	}
	//	<-ch
	//}
}
