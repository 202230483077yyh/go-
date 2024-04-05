package main

import "fmt"

// 定义接口类型
type Humaner interface {
	//方法只有声明，没有实现
	sayhi()
}

type student struct {
	name string
	id   int
}

func (tmp *student) sayhi() {
	fmt.Printf("student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type teacher struct {
	name string
	id   int
}

func (tmp teacher) sayhi() {
	fmt.Printf("teacher[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type mystr string

func (tmp *mystr) sayhi() {
	fmt.Println("hello , i am ", *tmp)
}

// 定义一个普通函数，函数参数为接口
// 只有一个函数，但有不同表现
func whosay(i Humaner) {
	i.sayhi()
}

func main() {
	//声明一个接口
	var i Humaner

	//实现接口函数的接收者类型变量，可以直接赋值给i
	s := student{"mike", 11}
	i = &s
	i.sayhi()

	t := teacher{"john", 88}
	i = t
	i.sayhi()

	var y mystr
	y = "yyh"
	i = &y
	i.sayhi()

	//多态
	whosay(&s)
	whosay(t)
	whosay(&y)

	x := make([]Humaner, 3)
	x[0] = &s
	x[1] = t
	x[2] = &y

	for _, val := range x {
		whosay(val)
	}

}
