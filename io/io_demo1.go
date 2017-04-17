package main

import (
	"io"
	"os"
	"fmt"
	"strings"
	"bufio"
	"bytes"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	var data []byte
	var err error

	// 从标准输入读取
	data, err = ReadFrom(os.Stdin, 11)
	fmt.Println(string(data), err)

	// 从普通文件读取，其中 file 是 os.File 的实例
	file, err := os.Open("README.md")
	data, err = ReadFrom(file, 9)
	fmt.Println(string(data), err)

	// 从字符串读取
	data, err = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(data), err)

	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)

	file, err = os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	n, err = file.WriteAt([]byte("Go语言中文网"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	fmt.Println("###########.gitignore###########")
	file, err = os.Open(".gitignore")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
	fileInfo, err := file.Stat()
	fmt.Printf("%+v, %+v\n", fileInfo, err)


	reader2 := bytes.NewReader([]byte("Go语言中文网"))
	reader2.WriteTo(os.Stdout)
	fmt.Println()

	reader3 := strings.NewReader("Go语言中文网")
	reader3.Seek(-6, os.SEEK_END)
	r, _, _ := reader3.ReadRune()
	fmt.Printf("%c\n", r)
}

