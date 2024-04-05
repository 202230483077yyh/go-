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

	m2[3] = "my"
	//map遍历(无序)
	for key, value := range m2 {
		fmt.Printf("%d ---->%s\n", key, value)
	}

	//如何判断一个key是否存在
	value, ok := m2[4]
	if ok == true {
		fmt.Println("m2[4]=", value)
	} else {
		fmt.Println("key不存在")
	}

}
