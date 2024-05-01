package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// 获取豆果网上的具体菜名的具体做法
func get_content(url string) (content string, err error) {
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
		content += string(buf[:n])
	}
	re := regexp.MustCompile(`<div class="step">\s*<h2.*?>(.*?)</h2>`)
	if re == nil {
		fmt.Println("regexp err")
		return
	}
	//2.取关键信息
	title := re.FindStringSubmatch(content)
	fmt.Println(title[1])

	stepInfoRegex := regexp.MustCompile(`<div class="stepinfo">\s*<p>(.*?)</p>(?s:(.*?))</div>`)

	// 使用正则表达式查找匹配项
	matches := stepInfoRegex.FindAllStringSubmatch(content, -1)

	// 遍历匹配结果
	for _, match := range matches {
		// 输出步骤编号和描述
		fmt.Println("步骤编号:", match[1])
		fmt.Println("步骤描述:", match[2])
		fmt.Println()
	}
	return
}

// 获取每一种类别的菜系下的具体菜urls
func get_urls(url string) (urls []string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	var content string

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body err = ", err)
			break
		}
		content += string(buf[:n])
	}

	//fmt.Println(content)

	re := regexp.MustCompile(`<li class="clearfix">\s*<a.*?href="([^"]+)".*?>`)
	tmp := re.FindAllStringSubmatch(content, -1)
	//fmt.Println(tmp)

	///cookbook/3300348.html
	//https://www.douguo.com/cookbook/3200196.html
	for _, data := range tmp {
		urls = append(urls, "https://www.douguo.com"+data[1])
	}
	return
}

// 获取指定图片url链接的图片，并下载本地
func get_picture(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// 获取具体菜的具体图片链接
func get_picture_url(url string) (urls []string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	var content string

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body err = ", err)
			break
		}
		content += string(buf[:n])
	}
	re := regexp.MustCompile(`<div class="stepcont clearfix">\s*<.*? href="(.*?)".*?>`)
	tmp := re.FindAllStringSubmatch(content, -1)
	for _, data := range tmp {
		urls = append(urls, data[1])
	}

	return
}

func main() {
	fmt.Println("正在爬取网页")
	url := "https://www.douguo.com/cookbook/3196958.html"
	urls, err1 := get_picture_url(url)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	filepath := "./img"
	str_chan := make(chan string)
	for id, data := range urls {
		//并发实现存在问题，解决方案：进程间通信，主进程一结束子进程自动扼杀！
		go func(id1 int, data1 string) {
			tmp := get_picture(data1, filepath+"/"+strconv.Itoa(id1)+".jpeg")
			if tmp != nil {
				fmt.Println("第", id1, "张图片获取失败", tmp)
			}
			str_chan <- strconv.Itoa(id1) + "爬取结束"
		}(id, data)
	}
	for i := 0; i < len(urls); i++ {
		tmp := <-str_chan
		fmt.Println(tmp)
	}
	return
}
