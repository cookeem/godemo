package main

import (
	"cookeem.com/ch2/common"
	"fmt"
)

func main() {
	i := common.Integer(1)
	var j common.Integer = 2
	fmt.Println("i.Add(j):", i.Add(j))

	fmt.Println("i.Add2(j):", i.Add2(&j))
	fmt.Println("i,j:", i, j)
}
