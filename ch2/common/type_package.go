package common

type Integer int

//类型方法，参数为值的方法
func (i Integer) Add(j Integer) (ret Integer) {
	ret = i + j
	return
}

//类型方法，参数为引用的方法
func (i *Integer) Add2(j *Integer) (ret Integer) {
	*i += 100
	*j += 10000
	ret = *i + *j
	return
}
