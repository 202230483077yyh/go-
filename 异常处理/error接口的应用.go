package main

import (
	"errors"
	"fmt"
)

func mydiv(a, b int) (res int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		res = a / b
	}
	return
}

func main() {
	res, err := mydiv(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("res = ", res)
	}

}
