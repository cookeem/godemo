package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			time.Sleep(time.Duration(3) * time.Second)
			ch <- 1
		}
	}()

	fmt.Println("begin!", time.Now())
	i := 0
LOOP: //for循环中的break label必须在for循环的前一句，goto的label则没有限制
	for {
		//select只会执行一次检测，外层必须加上for循环
		select {
		//设置超时机制
		case <-time.After(time.Duration(2) * time.Second):
			fmt.Println("timeout!", time.Now())
			i++
			//timeout超过3次退出
			if i > 3 {
				break LOOP
			}
		case <-ch:
			fmt.Println("<-ch", time.Now())
		}
	}
}
