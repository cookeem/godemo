package main

import (
	"errors"
	"fmt"
	"godemo/ch2/common"
)

type Rect struct {
	x, y          int
	width, height int
}

//不变更对象
func (rect Rect) Area() (area int, err error) {
	if rect.width < 1 || rect.height < 1 {
		err = errors.New("height or width must greater than 0")
		return
	}
	area = rect.width * rect.height
	return
}

//需要改变对象
func (rect *Rect) MoveToCenter() (err error) {
	if rect.width < 1 || rect.height < 1 {
		err = errors.New("height or width must greater than 0")
		return
	} else {
		rect.x = rect.width / 2
		rect.y = rect.height / 2
		return
	}
}

func (rect1 Rect) CompareArea(rect2 *Rect) (ret string, err error) {
	if rect1.width < 1 || rect1.height < 1 {
		err = errors.New("rect1 height or width must greater than 0")
		return
	}
	if rect2.width < 1 || rect2.height < 1 {
		err = errors.New("rect2 height or width must greater than 0")
		return
	}
	area1, _ := rect1.Area()
	area2, _ := rect2.Area()
	if area1 > area2 {
		ret = "bigger"
	} else if area1 == area2 {
		ret = "equal"
	} else {
		ret = "smaller"
	}
	return
}

func NewRect(x, y, width, height int) (rect Rect) {
	rect.x = x
	rect.y = y
	rect.height = height
	rect.width = width
	return rect
}

func NewRect2(x, y, width, height int) *Rect {
	return &Rect{x, y, width, height}
}

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

	arr1 := [3]int{1, 1, 1}
	arr2 := &arr1
	fmt.Println("arr1, *arr2:", arr1, *arr2)
	//数组值类型的修改引用
	arr2[0]++
	fmt.Println("arr1, *arr2:", arr1, *arr2)

	//slice,map,chan本身就是引用类型
	slice1 := []int{0, 0, 0}
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

	//结构体传值
	rect := Rect{-1, 1, 4, 2}
	if area, err := rect.Area(); err != nil {
		fmt.Println("rect.Area() error:", err)
	} else {
		fmt.Println("rect.Area():", area)
	}
	//结构体传引用，改变x和y的值
	if err := rect.MoveToCenter(); err != nil {
		fmt.Println("rect.MoveToCenter() error:", err)
	} else {
		fmt.Println("rect.MoveToCenter() rect.x, rect.y:", rect.x, rect.y)
	}

	rect1 := Rect{-1, 1, 3, 2}
	rect2 := &Rect{width: 5, height: 1}
	if ret, err := rect1.CompareArea(rect2); err != nil {
		fmt.Println("rect1.CompareArea(rect2) error:", err)
	} else {
		fmt.Println("rect1.CompareArea(rect2) ret:", ret)
	}

	rect3 := NewRect(1, 2, 5, 6)
	fmt.Println("rect3:", rect3)

	rect4 := NewRect2(0, 1, 4, 5)
	fmt.Println("*rect4:", *rect4)

	rect5 := new(Rect)
	fmt.Printf("rect5: %#v\n", rect5)

	fmt.Printf("str: %+q\n", "你好吗")
}
