package main

import (
	"fmt"
	"time"
)

func main() {
	// 简单使用 channel 的例子
	varChan1 := make(chan string)
	go func() {
		varChan1 <- "I am a string."
		close(varChan1)
	}()
	fmt.Println("Got msg: ", <-varChan1)
	fmt.Println("##########################")

	// 创建 channel 类型对象时设置了 buffer 值，buffer=0表示无缓冲，会引起写入阻塞
	varChan2 := make(chan string, 3)
	go func() {
		varChan2 <- "input 1"
		varChan2 <- "input 2"
		//只要是带有缓冲的channel，这句话都会最先执行，因为前边的3个语句是非阻塞的
		fmt.Println("I will be before all inputs")
		varChan2 <- "input 3"
		varChan2 <- "input 4"
		varChan2 <- "input 5"
		//建议goroutine关闭的时候，必须关闭channel，否则goroutine外读取就会出异常
		defer close(varChan2)
	}()
	//以死循环的方式监听channel，直到channel关闭则退出死循环
	for msg := range varChan2 {
		fmt.Println("Got msg:", msg, time.Now())
	}
	//检测channel是否已经关闭，否则循环读取，使用结果与上边的方式一致
	for {
		msg, ok := <-varChan2
		if ok {
			fmt.Println("Got msg: ", msg)
		} else {
			break
		}
	}
	fmt.Println("##########################")

	// channel 中的同步机制
	varChan3 := make(chan bool, 1)
	go func(varChan chan bool) {
		fmt.Println("begin to execute", time.Now())
		// do something
		time.Sleep(time.Second * 2)
		fmt.Println("end", time.Now())
		varChan <- true
		close(varChan3)
	}(varChan3)
	//这里会等待goroutine中channel的写入
	fmt.Println("Got msg: ", <-varChan3, time.Now())
	fmt.Println("##########################")

	// 单向 channel 类型对象的使用
	inChan := make(chan string, 1)
	outChan := make(chan string, 1)
	inChan <- "Input str"
	anonymousFunc := func(varInChan chan<- string, varOutChan <-chan string) {
		varInChan <- <-varOutChan
	}
	anonymousFunc(outChan, inChan)
	fmt.Println("Got msg: ", <-outChan)
	fmt.Println("##########################")

	// select的使用，记住select只是单次顺序判断，最先发生的情况就会执行，一般结合for使用
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1, time.Now())
		case msg2 := <-c2:
			fmt.Println("received", msg2, time.Now())
		}
	}
	fmt.Println("##########################")

	// select与超时检测
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "result 1"
	}()
	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	c4 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c4 <- "result 2"
	}()
	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
	fmt.Println("##########################")

	// select语句中default的优先级
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "msg"
		signals <- true
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg, time.Now())
	//因为有default存在，这个肯定最先执行，如果注释掉就会执行前一个语句
	default:
		fmt.Println("no message received", time.Now())
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg, time.Now())
	//因为有default存在，这个肯定最先执行，如果注释掉就会执行前一个语句
	default:
		fmt.Println("no message sent", time.Now())
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg, time.Now())
	case sig := <-signals:
		fmt.Println("received signal", sig, time.Now())
	//因为有default存在，这个肯定最先执行，如果注释掉就会执行前一个语句
	default:
		fmt.Println("no activity", time.Now())
	}
	fmt.Println("##########################")

	// 关闭channel，建议在goroutine结束的时候关闭channel，channel关闭后依然可以读取，但是写入会出异常
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//以死循环的方式，等待接收goroutine外发送的jobs。如果jobs已经关闭，那么退出goroutine
			j, more := <-jobs
			if more {
				fmt.Println("received job", j, time.Now())
			} else {
				fmt.Println("received all jobs, wait for 1 second to done", time.Now())
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j, time.Now())
		time.Sleep(1 * time.Second)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	time.Sleep(1 * time.Second)
	<-done
	fmt.Println("done", time.Now())
	fmt.Println("##########################")

	// for range方式循环读取buffer channel直到channel关闭
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
	fmt.Println("##########################")

	//jobs为单向发送channel，results为单向接收channel。
	numOfJobs := 12
	numOfWorkers := 4
	//jobs必须设置channel缓冲，否则任务多次写入channel会导致异常，最好设置缓冲大于任务书
	jobs2 := make(chan int, numOfJobs)
	results2 := make(chan int, numOfJobs)
	//启动numOfWorkers个worker
	for w := 1; w <= numOfWorkers; w++ {
		//每个worker接收jobs，输出results
		go worker(w, jobs2, results2)
	}
	time.Sleep(time.Second)
	//把numOfJobs个job分配给3个worker
	for j := 1; j <= numOfJobs; j++ {
		jobs2 <- j
	}
	//在jobs写入的程序段进行channel关闭
	close(jobs2)
	for a := 1; a <= numOfJobs; a++ {
		<-results2
	}
	fmt.Println("##########################")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	//在jobs关闭前不断从jobs中取数据
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j, time.Now())
		results <- j * 2
	}
}
