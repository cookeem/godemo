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

	//定义Man类型的变量man
	var man common.Man

	//men能存储Student
	man = &mike
	fmt.Println("This is Mike, a Student:")
	man.SayHi()
	man.Sing("November rain")

	//men也能存储Employee
	man = &tom
	fmt.Println("This is tom, an Employee:")
	man.SayHi()
	man.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	men := make([]common.Man, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	men[0], men[1], men[2] = &paul, &sam, &mike

	for _, value := range men {
		value.SayHi()
	}

	//接口的类型查询
	//man.(common.Employee)会失败，因为Employee有自己的SayHi方法
	if man2, ok := man.(common.Student); ok {
		fmt.Printf("%v is a Student\n", man2)
	} else {
		fmt.Printf("%v is not a Student\n", man)
	}
}
