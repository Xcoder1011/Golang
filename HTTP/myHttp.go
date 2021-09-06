package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	webServer()

	// request()
}

/// HTTP Server

func webServer() {

	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
	http.HandleFunc("/userInfo", userInfo)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("userInfo: wushangkun"))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Golang")
}

/// HTTP request

func request() {

	// url := "http://localhost:8000/"
	url := "http://www.baidu.com/"
	method := "GET"

	client := &http.Client{}

	// 创建一个http请求
	// 使用 strings.NewReader() 创建一个字符串的读取器，
	req, err := http.NewRequest(method, url, strings.NewReader("key=value"))

	// 发现错误就打印并退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	// 为标头添加信息
	req.Header.Add("User-Agent", "myClient")
	// 开始请求
	resp, err := client.Do(req)
	// 处理请求的错误
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
