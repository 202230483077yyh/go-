package main

import "fmt"

func main() {
	var ch byte    //'',不包括转义字符！
	var str string //""	字符串都是隐藏了一个'\0'

	ch = 'a'
	str = "abd"
	fmt.Printf("ch= %c\n str=%s\n", ch, str)
	fmt.Printf("str[0]= %c\n", str[0])

}
