package main

import (
	"cookeem.com/ch2/common"
	"fmt"
)

func main() {
	mike := common.Student{Human: common.Human{Name: "Mike", Age: 25, Phone: "222-222-XXX"}, School: "MIT", Loan: 0.00}
	paul := common.Student{Human: common.Human{Name: "Paul", Age: 26, Phone: "111-222-XXX"}, School: "Harvard", Loan: 100}
	sam := common.Employee{Human: common.Human{Name: "Sam", Age: 36, Phone: "444-222-XXX"}, Company: "Golang Inc.", Money: 1000}
	tom := common.Employee{Human: common.Human{Name: "Tom", Age: 37, Phone: "222-444-XXX"}, Company: "Things Ltd.", Money: 5000}

	tom.Work()

	//定义Men类型的变量i
	var i common.Men

	//i能存储Student
	i = &mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = &tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]common.Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = &paul, &sam, &mike

	for _, value := range x {
		value.SayHi()
	}

	fmt.Println("mike format println:", mike)
}
