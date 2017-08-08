package main

import (
	"fmt"
	"time"
)

func main() {
	// channel传递
	cht1 := make(chan int)
	cht2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now(), "cht1")
		cht1 <- 1
		return
	}()
	go func() {
		i := <-cht1
		time.Sleep(time.Second * 1)
		i++
		fmt.Println(time.Now(), "cht2", i)
		cht2 <- i
		return
	}()
	i := <-cht2
	fmt.Println(time.Now(), "cht2 finish:", i)

}
