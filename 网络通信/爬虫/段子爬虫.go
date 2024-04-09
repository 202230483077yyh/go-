package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Storefile(i int, filecontent, filetitle []string) {
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	defer f.Close()

	n := len(filetitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString("标题： " + filetitle[i])
		f.WriteString("\n")
		//写内容
		f.WriteString(filecontent[i])
		//分割
		f.WriteString("\n-----------------------------------------------\n")
	}
}

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		res += string(buf[:n])
	}
	return
}

func SpiderOneJoy(url string) (title, content string, err error) {
	res, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}

	//取解析式
	re1 := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err1")
		return
	}

	re2 := regexp.MustCompile(`<p>(?s:(.*?))</p>`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}

	//取内容
	tmpTitle := re1.FindAllStringSubmatch(res, -1)
	tmpContent := re2.FindAllStringSubmatch(res, -1)
	for _, data := range tmpTitle {
		title = data[0]
		//调整格式常用手段
		title = strings.Replace(title, "\t", "", -1)
		title = strings.Replace(title, "<title>", "", -1)
		title = strings.Replace(title, "</title>", "", -1)
		title = strings.Replace(title, "_捧腹网", "", -1)
		// title=strings.Replace(title,"\n","",-1)
		// title=strings.Replace(title,"\r","",-1)
		// title=strings.Replace(title," ","",-1)
	}

	for _, data := range tmpContent {
		//data[0]、[1]内容一样，不过[0]多了<p></p>
		// fmt.Println("mycontent1", data[0])
		//fmt.Println("mycontent2", data[1])
		content += data[1]
		content += "\n"
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, " ", "", -1)
		content = strings.Replace(content, "\r", "", -1)
	}

	return
}

func SpiderPage(i int) {
	//view-source:https://www.pengfu.net/e/tags/index.php?page=0&tagid=1&line=25&tempid=1
	//过滤标签：<h3 class="blogtitle"><a href=
	url := "https://www.pengfu.net/e/tags/index.php?page=" + strconv.Itoa(i) + "&tagid=1&line=25&tempid=1"

	fmt.Printf("正在爬取第%d个网页%s\n", i, url)

	//开始爬取页面全部内容，再过滤
	res, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//fmt.Println("r = ", res)
	//过滤关键信息
	//1.解析表达式
	re := regexp.MustCompile(`<h3 class="blogtitle"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp err")
		return
	}
	//2.取关键信息
	joyurls := re.FindAllStringSubmatch(res, -1)
	//fmt.Println(joyurls)

	filetitle := make([]string, 0)
	filecontent := make([]string, 0)

	for _, data := range joyurls {
		//fmt.Println("url = ", data[1])
		//开始爬取每一个笑话
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		// fmt.Println("title = ", title)
		// fmt.Println("content = ", content)

		//追加内容
		filetitle = append(filetitle, title)
		filecontent = append(filecontent, content)
	}

	//把内容写入到文件
	Storefile(i, filecontent, filetitle)

}

func DoWork(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)
	for i := start; i <= end; i++ {
		SpiderPage(i)
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
