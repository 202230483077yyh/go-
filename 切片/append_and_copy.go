package main

import (
	"fmt"
)

func main() {
	s := []int{}
	oldcap := cap(s)
	//以两倍的形式动态扩容
	for i := 0; i < 5; i++ {
		s = append(s, i)
		if newcap := cap(s); newcap > oldcap {
			fmt.Printf("oldcap %d --> newcap %d\n", oldcap, newcap)
			oldcap = newcap
		}
	}

	srcq := []int{6, 6, 6, 6, 6}
	desq := []int{1, 2}

	//copy不会改变原切片的容量长度，只会改变已有位置上的数
	copy(desq, srcq)
	fmt.Println(desq)
	//fmt.Println(srcq)

}
