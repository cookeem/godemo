package common

import (
	"errors"
	"strconv"
)

//多返回值函数
func Add(x, y int) (ret int, err error) {
	if x <= 0 || y <= 0 {
		err = errors.New("params must > 0!")
		return
	} else {
		ret = x + y
		return
	}
}

//不定参数函数
func Sum(args ...int) (ret int, err error){
	okParams := true
	for _, arg := range args {
		if (arg < 1) {
			okParams = false
			break
		}
	}
	if okParams {
		for _, arg := range args {
			ret += arg
		}
		return
	} else {
		err = errors.New("params must > 0!!")
		return
	}
}

//任意类型的不定参数函数
func Join(args ...interface{}) (ret string){
	for _, arg := range args {
		switch arg.(type) {
		case rune:
			//interface{} 类型推断
			v, ok := arg.(rune)
			if ok {
				ret += "rune:" + string(v) + ", "
			}
		case int:
			v, ok := arg.(int)
			if ok {
				ret += "int:" + strconv.Itoa(v) + ", "
			}
		case string:
			v, ok := arg.(string)
			if ok {
				ret += "string:" + v + ", "
			}
		default:
			ret += "unknown" + ", "
		}
	}
	return
}

const Pi = 3.14

