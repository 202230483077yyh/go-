package main

import (
	"fmt"
	"runtime"
)

func main() {
	//返回之前的 核数，并指定新的核数
	n := runtime.GOMAXPROCS(10)
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)

		fmt.Print(2)
	}
}
