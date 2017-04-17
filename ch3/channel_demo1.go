package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("@@@@@@@@@@@@@@@@@@", time.Now())
	c := make(chan bool)
	go func() {
		//必须在goroutine中激活channel，否则在goroutine外读取会出异常
		c <- true
		fmt.Println("Doing something…", time.Now())
	}()
	time.Sleep(time.Duration(1) * time.Second)
	<-c
	fmt.Println("Done!", time.Now())
	time.Sleep(time.Duration(2) * time.Second)

	fmt.Println("####################", time.Now())
	ch1 := make(chan int)
	//开启一个循环五次就会退出的goroutine
	go func() {
		//必须关闭ch1，否则gouroutine退出后，从ch1读取数据会panic
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(time.Duration(1) * time.Second)
			fmt.Println("ch1 <-", i, time.Now())
		}
	}()

	ch2 := make(chan int)
	//开启一个不会停止的goroutine
	go func() {
		i := 0
		for {
			fmt.Println("ch2 <-", i, time.Now())
			//在goroutine外从ch2接收消息，那么在goroutine内向ch2发送消息，此处会阻塞
			ch2 <- i
			time.Sleep(time.Duration(2) * time.Second)
			i++
		}
	}()
	//不断从goroutine的channel接收数据
	for {
		//ch1在close后依然可以读取，但是不能写入，可以同时判断ch1是否已经关闭
		if ret, ok := <-ch1; ok {
			fmt.Println("ch1 alive:", ret)
		} else {
			fmt.Println("ch1 close!")
		}
		//在goroutine内向ch2发送消息，那么在goroutine外从ch2接收消息，此处会阻塞
		v := <-ch2
		fmt.Println("ch2 read:", v, time.Now())
		//无缓冲channel情况下，用死循环的方式读取ch2
		//for v := range ch2 {
		//	fmt.Println("ch2 read:", v, time.Now())
		//}
	}
}
