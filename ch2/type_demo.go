package main

import (
	"cookeem.com/ch2/common"
	"fmt"
)

func main() {
	i := common.Integer(1)
	var j common.Integer = 2
	fmt.Println("i.Add(j):", i.Add(j))

	//值和引用
	fmt.Println("i.Add2(j):", i.Add2(&j))
	fmt.Println("i,j:", i, j)

	a := 0
	b := &a
	fmt.Println("a, *b:", a, *b)
	//简单值类型的修改引用
	*b++
	fmt.Println("a, *b:", a, *b)

	arr1 := [3]int { 1, 1, 1}
	arr2 := &arr1
	fmt.Println("arr1, *arr2:", arr1, *arr2)
	//数组值类型的修改引用
	arr2[0]++
	fmt.Println("arr1, *arr2:", arr1, *arr2)

	//slice,map,chan本身就是引用类型
	slice1 := []int { 0, 0, 0 }
	slice2 := slice1
	fmt.Println("slice1, slice2:", slice1, slice2)
	slice2[0] = 1
	fmt.Println("slice1, slice2:", slice1, slice2)

	//interface{}也是属于值类型
	var intf1 interface{} = 0
	intf2 := &intf1
	fmt.Println("intf1, *intf2:", intf1, *intf2)
	if i, ok := (*intf2).(int); ok {
		*intf2 = i + 1
	}
	fmt.Println("intf1, *intf2:", intf1, *intf2)
}
