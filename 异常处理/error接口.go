package main

import (
	"errors"
	"fmt"
)

func main() {
	errl := fmt.Errorf("%s", "this is normal err")
	fmt.Println("errl = ", errl)

	errl2 := errors.New("this is normal err2")
	fmt.Println("errl2 = ", errl2)
}
