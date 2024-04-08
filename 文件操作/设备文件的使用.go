package main

import (
	"fmt"
	"os"
)

func main() {
	//把输出关了，就不会有打印
	//os.Stdout.Close()
	fmt.Printf("are you ok?\n")

	//标准设备文件(os.Stdout)，默认已经打开，用户可以直接使用
	os.Stdout.WriteString("yes \n")

	//把输入关了，Scan就无效了
	//os.Stdin.Close()
	var a int
	fmt.Println("请输入a: ")
	fmt.Scan(&a)
	fmt.Println("a = ", a)

}
