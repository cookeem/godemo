package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("block for 1 second", time.Now())
	time.Sleep(time.Second * 1)
	//等待向c发送消息
	c <- sum
	fmt.Println("block end", time.Now())
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := make(chan int)

	//通过一个channel同时启动多个goroutine执行任务
	//通过控制goroutine内向channel发送消息，goroutine外从channel接收消息，实现多个goroutine的阻塞和并行。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y, time.Now())
}
