package main

import (
	"fmt"
	"net/http"
)

func main() {
	//记得要加上http
	resp, err := http.Get("http://localhost:8000/mydream")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)
	//Body是一个io指针，需要解指针
	fmt.Println("Body = ", resp.Body)

	var tmp string
	tmp = ""
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println(tmp)
}
