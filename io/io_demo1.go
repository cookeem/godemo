package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
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

	fmt.Println("从标准输入读取数据")
	data, err = ReadFrom(os.Stdin, 11)
	fmt.Println(string(data), err)

	fmt.Println("##从普通文件读取数据")
	file, err := os.Open("cookeem.com.iml")
	if err != nil {
		fmt.Println("##从普通文件读取数据错误：", err)
		return
	}
	data, err = ReadFrom(file, 9)
	fmt.Println(string(data), err)
	defer file.Close()

	fmt.Println("##从字符串读取数据")
	data, err = ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(data), err)

	fmt.Println("##从字符串读取特定偏移量数据到bytes")
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 3)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		fmt.Println("从字符串读取特定偏移量数据错误：", err)
		return
	}
	fmt.Printf("%s, %d\n", p, n)

	fmt.Println("##并替换特定区域字符串，并写入文件")
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

	fmt.Println("##读取特定文件，并输出到屏幕")
	file, err = os.Open("cookeem.com.iml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
	fileInfo, err := file.Stat()
	fmt.Printf("%+v, %+v\n", fileInfo, err)

	fmt.Println("##字符串打印到屏幕")
	reader2 := bytes.NewReader([]byte("Go语言中文网\n"))
	reader2.WriteTo(os.Stdout)

	fmt.Println("##获取特定区域字符，倒数第x个字符")
	reader3 := strings.NewReader("Go语言中文网")
	reader3.Seek(-6, os.SEEK_END)
	r, _, _ := reader3.ReadRune()
	fmt.Printf("%c\n", r)

	fmt.Println("##压缩文件")
	dat, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fileOut, err := os.Create("README.gz")
	if err != nil {
		fmt.Println(err)
	}
	defer fileOut.Close()
	writerGz := gzip.NewWriter(fileOut)
	writerGz.Write(dat)
	defer writerGz.Flush()
	defer writerGz.Close()

	fmt.Println("##CopyBuffer")
	r1 := strings.NewReader("first reader\n")
	buf := make([]byte, 3)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	fmt.Println("##读取一个输入的字符")
	var ch byte
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err = buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	} else {
		fmt.Println("写入错误")
	}

	fmt.Println("##io.Copy方法")
	io.Copy(os.Stdout, strings.NewReader("测试Copy方法\n"))
	io.CopyN(os.Stdout, strings.NewReader("Go语言中文网"), 8)
	fmt.Println()

	fmt.Println("##io.ReadFull方法")
	str := "GO语言中文网"
	var b = make([]byte, len(str))
	n, err = io.ReadFull(strings.NewReader(str), b)
	if err != nil {
		return
	}
	fmt.Println(string(b), n)

	fmt.Println("##MultiReader使用，多个reader变成一个reader处理")
	readers := []io.Reader{
		strings.NewReader("from strings reader"),
		bytes.NewBufferString("from bytes buffer"),
	}
	readerMulit := io.MultiReader(readers...)
	data2 := make([]byte, 0, 1024)
	for err != io.EOF {
		tmp := make([]byte, 512)
		n, err = readerMulit.Read(tmp)
		if err == nil {
			data2 = append(data2, tmp[:n]...)
		} else {
			if err != io.EOF {
				panic(err)
			}
		}
	}
	fmt.Printf("%s\n", data2)

	fmt.Println("##MultiWriter使用，多个writer变成一个writer处理")
	file, err = os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	writerMulti := io.MultiWriter(writers...)
	writerMulti.Write([]byte("Go语言中文网\n"))

	fmt.Println("##TeeReader用法，有点类似stream")
	readerTee := io.TeeReader(strings.NewReader("Go语言中文网\n"), os.Stdout)
	readerTee.Read(make([]byte, 20))

	fmt.Println("##ReadAll用法")
	dat, err = ioutil.ReadAll(strings.NewReader("hello world"))
	if err != nil {
		fmt.Println("ioutil.ReadAll error:", err)
		return
	}
	fmt.Println(string(dat))

	fmt.Println("##ReadDir用法")
	fis, err := ioutil.ReadDir("ch1")
	for _, fi := range fis {
		fmt.Println(fi.Name(), fi.IsDir(), fi.Size(), fi.ModTime())
	}

	fmt.Println("##Scanner统计词频")
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println(count)

	fmt.Println("##NewScanner从标准输入中获取数据")
	scanner = bufio.NewScanner(strings.NewReader("hello\nworld\nhaijian\n"))
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	now := time.Now()
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	fmt.Println(year, int(month), day, hour, minute, second)
	t2 := time.Date(2010, 10, 20, 12, 24, 36, 512, time.Local)
	fmt.Println(t2)

}
