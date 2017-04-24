package main

import (
	"bufio"
	"fmt"
	"net/http"
	"io/ioutil"
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
		fmt.Println(string(line), isPrefix, err)
		fmt.Println("###########")
		line, isPrefix, err = buf.ReadLine()
		if err != nil {
			fmt.Println("end of file error:", err)
			break
		}
	}


	fmt.Println("## 使用bufio每20个字符读")
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
		fmt.Println(string(bytes), n, err)
		fmt.Println("###########")
		n, err = buf.Read(bytes)
		if err != nil {
			fmt.Println("end of file error:", err)
			break
		}
	}

}
