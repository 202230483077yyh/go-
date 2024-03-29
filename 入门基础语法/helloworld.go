// go语言以包作为管理单位，每个文件必须先声明包
// 程序要运行必须有一个main包！！
package main

import "fmt"

// 入口函数一个工程（文件夹）只有一个！
func main() { //左括号必须与函数名同行
	//调用函数需要导入包，Println()自动换行
	fmt.Println("hello world!")
}
