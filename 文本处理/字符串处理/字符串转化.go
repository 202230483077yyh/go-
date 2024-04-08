package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转化为字符串后，追加到字符数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)

	//第二个数为要追加的数，第三个数为指定10进制的方式追加
	slice = strconv.AppendInt(slice, 123, 10)
	slice = strconv.AppendQuote(slice, "avdd")

	fmt.Println("slice = ", string(slice))

	//其他类型转化为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)

	//'f'指打印格式，以小数方式，-1指小数点位数（紧缩模式），64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str = ", str)

	//整型转字符串
	str = strconv.Itoa(666)
	fmt.Println("str = ", str)

	//字符串转其他类型
	//1.字符转bool
	var flag bool
	var err error
	flag, err = strconv.ParseBool("alse")
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("flag = ", flag)
	}

	//2.字符串转整型
	a, _ := strconv.Atoi("567")
	fmt.Println("a = ", a)

}
