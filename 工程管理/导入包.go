package main

//1.包含多个包	用圆括号
// import (
// 	"fmt"
// 	"os"
// )

//2.给包起别名
import io "fmt"
import "os"

//3.	.操作
//import ."fmt"
//可以不写包名，直接调用函数（易导致命名冲突）

//4.忽略此包
//import _ "fmt"

func main() {
	io.Println("OS.ARG = ", os.Args)
}
