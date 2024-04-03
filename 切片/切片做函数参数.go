package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(v []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		v[i] = rand.Intn(100)
	}
}

func BubbleSort(v []int) {
	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-1-i; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}
}

func main() {
	s := make([]int, 10)

	InitData(s)
	fmt.Println("排序前： ", s)

	BubbleSort(s)
	fmt.Println("排序后： ", s)
}
