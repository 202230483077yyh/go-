package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body err = ", err)
			break
		}
		res += string(buf[:n])
	}
	return
}

func Spider(i int, page chan<- int) {

	fmt.Printf("正在爬第%d个网页", i)
	//1.拼接url
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa(50*(i-1))
	//2.爬取网页所有内容
	res, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	//3.把内容写入到文件
	filename := strconv.Itoa(i) + ".html"
	f, ferr := os.Create(filename)
	if ferr != nil {
		fmt.Println("ferr = ", ferr)
		return
	}

	f.WriteString(res)
	f.Close()

	page <- i
}

func DoWork(start, end int) {
	page := make(chan int)
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)
	for i := start; i <= end; i++ {
		go Spider(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}

}

func main() {
	var start, end int
	fmt.Println("请输入起始页（>=1）：")
	fmt.Scan(&start)

	fmt.Println("请输入终止页：")
	fmt.Scan(&end)

	DoWork(start, end)

}
