package main

import "fmt"

func main() {
	//if语句可以支持一个初始化语句;判断语句
	//if语句没有小括号
	if a := 10; a == 10 {
		fmt.Printf("OK\n")
	}

	//注意使用if_else if_else保证紧挨着上一个右括号
	b := 10
	if b == 10 {
		fmt.Printf("OK\n")
	} else if b > 10 {
		fmt.Printf("b>10\n")
	} else {
		fmt.Printf("b<10")
	}

	//switch语句，也可以支持初始化语句
	num := 10
	switch num {
	case 1:
		fmt.Printf("one floor\n")
		break //break也可以不写，默认包含
		//fallthrough，取消默认写break，即不跳出switch
	case 2:
		fmt.Printf("two floor\n")
		break
	default:
		fmt.Printf("invalid\n")
	}

	score := 89
	switch { //可以不指定变量
	case score >= 90: //在case之后放判断条件
		fmt.Printf("yyds\n")

	case score < 90:
		fmt.Printf("laji\n")
	}

}
