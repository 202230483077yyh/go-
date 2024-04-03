package main

import "fmt"

func main() {
	//对map只有len，没有cap
	tmp := map[int]string{
		100: "abcd",
		101: "林",
	}
	tmp[222] = "niuma"
	fmt.Println(tmp)
	fmt.Println(len(tmp))

	//同样可以make创建,可以指定长度（可以不用），长度后期也会自动扩容
	m2 := make(map[int]string, 10)
	m2[1] = "mike"
	fmt.Println(m2)
	fmt.Println(len(m2))
}
