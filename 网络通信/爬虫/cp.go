package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Dish struct {
	name         string
	operation    []string
	picture_urls []string
	tag          string
}

// 获取豆果网上的具体菜名的具体做法
func get_content(url string) (name string, operation []string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 1024*4)
	var content string
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
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
	if len(title) < 1 {
		fmt.Println("错误的url", url)
		return
	}
	name = title[1]

	stepInfoRegex := regexp.MustCompile(`<div class="stepinfo">\s*<p>(.*?)</p>(?s:(.*?))</div>`)

	// 使用正则表达式查找匹配项
	matches := stepInfoRegex.FindAllStringSubmatch(content, -1)

	// 遍历匹配结果
	for _, match := range matches {
		// 输出步骤编号和描述
		// fmt.Println("步骤编号:", match[1])
		// fmt.Println("步骤描述:", match[2])

		//调整格式
		stmp := strings.Trim(match[2], "\t")
		stmp = strings.Trim(stmp, " ")
		stmp = strings.Trim(stmp, "\n")
		stmp = strings.Trim(stmp, "\r")
		operation = append(operation, stmp)
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
		n, _ := resp.Body.Read(buf)
		if n == 0 {
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
		n, _ := resp.Body.Read(buf)
		if n == 0 {
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

// 获取指定url的所有图片并保存到本地
func download_pictures(url string) {
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
		<-str_chan
		//fmt.Println(tmp)
	}
	return
}

// 获取具体菜的封装类
func get_class(url string) (res Dish, err error) {
	res.name, res.operation, err = get_content(url)
	if err != nil {
		return
	}
	res.picture_urls, err = get_picture_url(url)
	if err != nil {
		return
	}
	return
}

func Print_Dish(res Dish) {
	fmt.Println(res.name)
	fmt.Println("具体步骤如下：")
	for _, data := range res.operation {
		fmt.Println(data)
	}
	fmt.Println("具体图片链接如下：")
	for _, data := range res.picture_urls {
		fmt.Println(data)
	}
	return
}

// 将类转化为项
func to_db(res Dish, db *sql.DB) {
	m1 := make(map[int]string, 1)
	for id, data := range res.operation {
		m1[id] = data
	}

	m2 := make(map[int]string, 1)
	for id, data := range res.picture_urls {
		m2[id] = data
	}
	m11, _ := json.MarshalIndent(m1, "", " ")
	m21, _ := json.MarshalIndent(m2, "", " ")

	_, err := db.Exec(`insert into dish(name,operation,picture_urls) values(?,?,?)`, res.name, m11, m21)
	if err != nil {
		fmt.Println("插入失败：", err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建新的数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS caipu")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database created successfully.")

	_, err = db.Exec(`use caipu`)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("切换caipu数据库")

	//往数据库里新建表
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dish(name VARCHAR(255),operation json,picture_urls json)`)
	if err != nil {
		panic(err.Error())
	}

	url0 := "https://www.douguo.com/caipu/海鲜"
	urls, err0 := get_urls(url0)
	if err0 != nil {
		fmt.Println("获取菜系各个菜的url失败：", err0)
		return
	}

	for _, url := range urls {
		res, err := get_class(url)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		// Print_Dish(res)
		to_db(res, db)
		time.Sleep(2 * time.Second)
	}

}

// func main() {
// 	url := "https://www.douguo.com/cookbook/2506854.html"
// 	name, operation, err := get_content(url)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return
// 	}
// 	fmt.Println(name)
// 	fmt.Println(operation)
// }

// func main() {
// 	url0 := "https://www.douguo.com/caipu/海鲜"
// 	urls, err0 := get_urls(url0)
// 	if err0 != nil {
// 		fmt.Println("获取菜系各个菜的url失败：", err0)
// 		return
// 	}

// 	for _, url := range urls {
// 		fmt.Println(url)
// 	}
// 	return
// }
