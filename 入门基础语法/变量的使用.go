package main

import "fmt"

//导入包后，必须用，否则报错

func main() {
	//变量，程序运行期间可以改变的量
	//在同一个{}变量声明唯一
	//1.声明格式：	var 变量名 类型	(变量声明了，必须使用(赋值也算)，默认为0)
	var a int
	//var b, c int
	fmt.Println("a=", a)

	//2.变量的初始化，声明变量时赋值
	//var d int = 10
	//d = 30

	//3.自动推导类型，必须初始化，以此确定类型
	c := 30
	//Printf格式化输出并不会自动换行，Println一段段输出会自动换行
	fmt.Printf("c type is %T\n", c)

	// 不同类型变量的声明
	var (
		a1 int
		b1 float64
	)
	a1, b1 = 10, 3.14
	fmt.Println("a1=", a1, "b1=", b1)

}
