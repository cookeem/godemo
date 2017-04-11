package main

import (
	"fmt"
	"os"
)

func main() {
	//if条件语句
	args := os.Args

	if i, argsSize := 1, len(args); argsSize - i == 0 {
		fmt.Println("no args")
	} else if argsSize > 1 && argsSize < 3 {
		fmt.Println("argsSize > 1 && argsSize < 3")
	} else {
		fmt.Println("argsSize >= 3", argsSize)
	}

	//switch语句
	j := 1
	switch j {
	case 0:
		fmt.Println("j == 0")
	case 1:
		fmt.Println("j == 1")
		fallthrough
	case 2:
		fmt.Println("j == 2")
	default:
		fmt.Println("j default")
	}

	k := 1
	switch {
	case k == 0:
		fmt.Println("k == 0")
	case k == 1:
		fmt.Println("k == 1")
		fallthrough
	case k == 2:
		fmt.Println("k == 2")
	default:
		fmt.Println("k default")
	}

	//循环
	for i, v := range args {
		if i > 0 {
			fmt.Println(i, v)
		}
	}

	//while模式
	i := 0
	for {
		if i > 5 {
			break
		}
		fmt.Println("loop1 i:", i)
		i++
	}

	for i := 0;; {
		if i > 5 {
			break
		}
		fmt.Println("loop2 i:", i)
		i++
	}
}
