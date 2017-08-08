package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

//自定义错误
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

func Stat(name string) (err error) {
	var stat syscall.Stat_t
	err = syscall.Stat(name, &stat)
	if err != nil {
		err = PathError{"stat", name, err}
		return
	}
	return
}

//延迟处理错误
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("os.Open(", src, ") err:", err)
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("os.Create(", dst, ") err:", err)
		return
	}
	defer dstFile.Close()
	w, err = io.Copy(dstFile, srcFile)
	return
}

func SimplePanicRecover() {
	//defer recover必须在panic之前
	defer func() {
		if err := recover(); err != nil {
			e, ok := err.(error)
			if ok {
				fmt.Println("e.Error() is:", e.Error())
			}
			fmt.Println("Panic info is: ", err)
		}
	}()
	i := 1
	fmt.Println(1 / (i - 1))
	//该语句不会执行，因为上一个语句已经panic被捕获
	fmt.Println("###panic finish")
}

// 当 defer 中也调用了 panic 函数时，最后被调用的 panic 函数的参数会被后面的 recover 函数获取到
// 一个函数中可以定义多个 defer 函数，按照 FILO 的规则执行
func MultiPanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("1 Panic info is: ", err)
		}
	}()
	defer func() {
		panic("1 MultiPanicRecover defer inner panic")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("2 Panic info is: ", err)
		}
	}()
	panic("2 MultiPanicRecover function panic-ed!")
}

// recover 函数只有在 defer 函数中被直接调用的时候才可以获取 panic 的参数
func RecoverPlaceTest() {
	// 下面一行代码中 recover 函数会返回 nil，但也不影响程序运行
	defer recover()
	// recover 函数返回 nil
	defer fmt.Println("RecoverPlaceTest recover() is: ", recover())
	defer func() {
		func() {
			// 由于不是在 defer 调用函数中直接调用 recover 函数，recover 函数会返回 nil
			if err := recover(); err != nil {
				fmt.Println("RecoverPlaceTest Panic info is: ", err)
			}
		}()

	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RecoverPlaceTest Panic info is: ", err)
		}
	}()
	panic("RecoverPlaceTest function panic-ed!")
}

// 如果函数没有 panic，调用 recover 函数不会获取到任何信息，也不会影响当前进程。
func NoPanicButHasRecover() {
	if err := recover(); err != nil {
		fmt.Println("NoPanicButHasRecover Panic info is: ", err)
	} else {
		fmt.Println("NoPanicButHasRecover Panic info is: ", err)
	}
}

// 定义一个调用 recover 函数的函数
func CallRecover() {
	if err := recover(); err != nil {
		fmt.Println("CallRecover Panic info is: ", err)
	}
}

// 定义个函数，在其中 defer 另一个调用了 recover 函数的函数
func RecoverInOutterFunc() {
	defer CallRecover()
	panic("RecoverInOutterFunc function panic-ed!")
}

func main() {
	err := Stat("a.txt")
	if err != nil {
		fmt.Println("err:", err.Error())
	}

	CopyFile("test.txt", "test.txt.bak")

	SimplePanicRecover()
	MultiPanicRecover()
	RecoverPlaceTest()
	NoPanicButHasRecover()
	RecoverInOutterFunc()

	fmt.Println("adsfa" + "\n" +
		"sdfasdf" + "\n" +
		"dsasdfasdf")

	//defer执行顺序，FILO，first in last out
	defer func() {
		fmt.Println("defer func 1")
	}()

	defer func() {
		fmt.Println("defer func 2")
	}()

	defer func() {
		fmt.Println("defer func 3")
	}()
}
