package main

import (
	"errors"
	"fmt"
	"strings"
)

func GetName(name string) (firstName, midName, lastName string, err error) {
	arr := strings.Split(name, " ")
	if len(arr) == 3 {
		firstName = arr[0]
		midName = arr[1]
		lastName = arr[2]
	} else {
		err = errors.New("Name must include 3 segment split by space")
	}
	return
}

func main() {
	//变量和常量
	const (
		Sunday = iota
		Monday
		Tuesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Thursday, Friday, Saturday)

	var (
		i int     = 1
		j float32 = 2
		k float32 = 1
	)
	l, m := j/k, k/j
	if j == l {
		fmt.Println("i:", i, "j:", j, "l:", l, "m:", m)
	} else {
		fmt.Println("not match")
	}
	fmt.Println("函数调用1")
	fn, mn, ln, err := GetName("zeng hai jian")
	if err == nil {
		fmt.Println(fn, mn, ln)
	} else {
		fmt.Println("err: ", err)
	}
	fmt.Println("函数调用2")
	fn2, mn, ln, err := GetName("wu wen jing")
	if err == nil {
		fmt.Println(fn2, mn, ln)
	} else {
		fmt.Println("err: ", err)
	}

	//字符串
	s1 := "曾海😂"
	fmt.Println("first char:", string(s1[0]), "string len:", len(s1))
	fmt.Println(string(s1[:3]))
	fmt.Println(string(s1[len(s1)-4:]))

	fmt.Println(strings.Contains(s1, "曾"))
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}
	fmt.Println("############")
	for i, v := range s1 {
		fmt.Println(i, string(v))
	}

	fmt.Println("$$$$$$$$$$$$")
	f := func(c rune) bool {
		fmt.Println(string(c))
		if c > 10 {
			return true
		} else {
			return false
		}
	}
	println(strings.IndexFunc(s1, f))

	//数组
	type Point struct{ x, y int }
	arr := [5]struct{ a, b int }{{1, 2}, {3, 4}}
	for i, v := range arr {
		fmt.Println("arr:", i, v)
	}
	arr2 := [5]int{1, 2, 3}
	arr2[4] = 5
	for i := 0; i < len(arr2); i++ {
		fmt.Println("arr2:", i, arr2[i])
	}
	arr3 := [3][2]int{}
	fmt.Println("arr3:", arr3)

	//切片
	arr4 := [5]Point{{}, {}, {3, 3}, {}, {5, 5}}
	slice1 := make([]Point, 5, 10)
	slice1 = arr4[2:]
	fmt.Println("slice1:", slice1, "len(slice1):", len(slice1), "cap(slice1):", cap(slice1))
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 5)
	slice3 := []int{6, 7}
	slice2 = append(slice2, slice3...)
	fmt.Println("slice2:", slice2, "len(slice2):", len(slice2), "cap(slice2):", cap(slice2))
	copy(slice2, slice3)
	fmt.Println("slice2:", slice2, "len(slice2):", len(slice2), "cap(slice2):", cap(slice2))

	//map
	type Person struct {
		Name string
		Age  int
	}
	map1 := map[string]Person{
		"NO1": Person{"haijian", 38},
		"NO2": Person{"wenjing", 31},
	}
	fmt.Println("map1:", map1, "len(map1):", len(map1))
	for k, v := range map1 {
		fmt.Println(k, v.Name, v.Age)
	}

	map2 := make(map[string]string, 2)
	map2["k1"] = "v1"
	map2["k2"] = "v2"
	map2["k3"] = "v3"
	map2["k4"] = "v4"
	delete(map2, "k1")
	delete(map2, "k5")
	fmt.Println("map2:", map2, "len(map2):", len(map2))

	map2V, ok := map2["k3"]
	if ok {
		fmt.Println("map2V:", map2V)
	} else {
		fmt.Println("ok:", ok)
	}

	map2V2, ok := map2["k2"]
	if ok {
		fmt.Println("map2V2:", map2V2)
	} else {
		fmt.Println("ok:", ok)
	}

	s := fmt.Sprintf("map1: %v\nmap2: %v", map1, map2)
	fmt.Println(s)

}
