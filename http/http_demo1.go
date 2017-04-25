package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.baidu.com"
	var resp *http.Response
	var err error

	fmt.Println("## 使用ioutil全量读reader")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read all error:", err)
		return
	}
	fmt.Println(string(bytes))

	fmt.Println("## 使用bufio逐行读reader")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	buf := bufio.NewReader(resp.Body)
	if err != nil {
		fmt.Println("bufio.NewReader error:", err)
		return
	}
	for line, isPrefix, err := buf.ReadLine(); ; {
		fmt.Println(string(line), isPrefix)
		fmt.Println("###########")
		line, isPrefix, err = buf.ReadLine()
		if err != nil {
			fmt.Println("end of file error:", err)
			break
		}
	}

	fmt.Println("## 使用bufio读取，遇到20个字符停止读取")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	buf = bufio.NewReader(resp.Body)
	if err != nil {
		fmt.Println("bufio.NewReader error:", err)
		return
	}
	bytes = make([]byte, 20)
	for n, err := buf.Read(bytes); ; {
		fmt.Println(string(bytes), n)
		fmt.Println("###########")
		n, err = buf.Read(bytes)
		if err != nil {
			fmt.Println("end of file error:", err)
			break
		}
	}

	fmt.Println("## 使用bufio读取，遇到右尖括号停止读取")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("http get error:", err)
		return
	}
	defer resp.Body.Close()
	buf = bufio.NewReader(resp.Body)
	if err != nil {
		fmt.Println("bufio.NewReader error:", err)
		return
	}
	bytes = make([]byte, 20)
	for segment, err := buf.ReadSlice('>'); ; {
		fmt.Println(string(segment))
		fmt.Println("###########")
		segment, err = buf.ReadSlice('>')
		if err != nil {
			fmt.Println("end of file error:", err)
			break
		}
	}

}
