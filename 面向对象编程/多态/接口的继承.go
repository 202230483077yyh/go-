package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type personer interface { //超集
	Humaner
	sing(lrc string)
}

type student struct {
	id   int
	name string
}

func (tmp *student) sayhi() {
	fmt.Println("hello , i am ", tmp.name)
}

func (tmp *student) sing(lrc string) {
	fmt.Println("Stu在唱着歌", lrc)
}

func main() {

	var i personer
	s := student{2022, "yyh"}
	i = &s
	i.sayhi()
	i.sing("洛天依")

	//超集可以转化为子集，子集不可以转化为超集
	var i2 Humaner

	//i=i2
	i2 = i
	i2.sayhi()

}
