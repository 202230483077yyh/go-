package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func main() {
	s1 := student{1, "mike", 16}
	s2 := student{1, "mike", 16}
	s3 := student{2, "mike", 16}

	//同类型比较
	fmt.Println("s1==s2", s1 == s2)
	fmt.Println("s2==s2", s2 == s3)

	//同类型赋值
	var tmp student
	tmp = s3
	fmt.Println("tmp = ", tmp)

}
