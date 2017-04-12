package main

import (
	"cookeem.com/ch1/common"
	"fmt"
)

func main() {
	x, y := 1, 2
	ret, err := common.Add(x, y)
	if err != nil {
		fmt.Println("Add(", x, ",", y, ") err!", err)
	} else {
		fmt.Println("Add(", x, ",", y, "):", ret)
	}

	arr := []int{1, 2, 3}
	sum, err := common.Sum(arr...)
	if err != nil {
		fmt.Println("Sum(", arr, ") err!", err)
	} else {
		fmt.Println("Sum(", arr, "):", sum)
	}

	arr2 := make([]int, 5)
	sum, err = common.Sum(arr2...)
	if err != nil {
		fmt.Println("Sum(", arr2, ") err!", err)
	} else {
		fmt.Println("Sum(", arr2, "):", sum)
	}

	arr3 := []interface{}{int(10), "haha", rune(300)}
	joinStr := common.Join(arr3...)
	fmt.Println("common.Join(", arr3, ")", joinStr)

	fmt.Println("common.Pi:", common.Pi)

	//函数赋值给变量
	func1 := func(x, y int) (ret int) {
		ret = x * y
		return
	}
	fmt.Println("f(3, 4):", func1(3, 4))

	//匿名函数，高阶函数，函数的参数是函数
	func(i int, j int, f func(x, y int) (ret int)) {
		fmt.Println("func(5, 6, func1):", f(i, j))
	}(5, 6, func1)

}
