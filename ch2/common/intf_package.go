package common

import (
	"fmt"
	"strconv"
)

type Human struct {
	Name  string
	Age   int
	Phone string
}

type Student struct {
	Human  //匿名字段
	School string
	Loan   float32
}

type Employee struct {
	Human   //匿名字段
	Company string
	Money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//fmt.Println中定义了String()函数，如果Human也定义了String函数，那么调用fmt.Println的时候将会使用Human.String()函数进行打印
func (h Human) String() string {
	return "《" + h.Name + " - " + strconv.Itoa(h.Age) + " years -  ✆ " + h.Phone + "》"
}

//Employee重载Human的SayHi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name, e.Company, e.Phone)
}

//Employee实现Work方法
func (e *Employee) Work() {
	fmt.Printf("Hi, I am %s, I am working\n", e.Name)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}
