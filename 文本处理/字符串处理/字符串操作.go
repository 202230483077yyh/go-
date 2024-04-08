package main

import (
	"fmt"
	"strings"
)

func main() {
	//查看是否包含，返回true或者false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("gogoday", "hello"))

	//Joins组合
	//将字符串与字符串通过指定字符拼接
	s := []string{"abc", "deg", "ftr"}
	fmt.Println("s = ", s)
	buf := strings.Join(s, ",")
	fmt.Println("buf = ", buf)

	//Index	返回子串在字符串中的索引（若不存在返回-1）
	fmt.Println(strings.Index("abchello", "hello"))
	fmt.Println(strings.Index("abcde", "abd"))

	//Repeat 指定字符串重复拼接次数
	buf = strings.Repeat("go", 2)
	fmt.Println("buf = ", buf)

	//Split	按指定字符分割
	//元素放入切片中
	buf = "hello%bcd%mike"
	s2 := strings.Split(buf, "%")
	for i, data := range s2 {
		fmt.Println(i, " ", data)
	}

	//Trim 去掉两头的指定字符
	buf = strings.Trim("    are you ok?       ", " ")
	fmt.Println("buf=", buf)

	//Fields特殊化的Split，只去除空格
	//元素放入切片中
	s3 := strings.Fields(" are you ok ? ")
	for i, data := range s3 {
		fmt.Println(i, " ", data)
	}

}
