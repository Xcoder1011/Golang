package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	fmt.Println("Hello! wushangkun")

	// url := "http://localhost:8000/"
	url := "http://www.baidu.com/"

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
