package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string
	Age  int
}

type Website struct {
	url string
}

func (u User) String() (ret string) {
	ret = fmt.Sprintf("Hi I am %s, I am %d years old", u.Name, u.Age)
	return ret
}

func main() {
	user := User{Name: "haijian", Age: 39}
	website := Website{"https://cookeem.com"}

	fmt.Printf("%v, %+v, %#v, %T\n", user, website, website, website)

	arr := []string{"str1", "str2", "曾海剑"}
	fmt.Printf("%v\n", arr)

	str := "曾海剑"
	str2 := fmt.Sprintf("%+q", str)
	fmt.Printf("%+s\n", str2)

	fmt.Fprintf(os.Stdout, "%+s\n", str2)

	var args1, args2 string
	fmt.Scanln(&args1, &args2)
	fmt.Println(args1, args2)

	filePath := "/home/polaris/studygolang"
	sp := strings.Split(filePath, "/")
	for i, v := range sp {
		fmt.Println(i, v)
	}
}
