package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	//我们设置了2个路由，/和/hello，分别绑定 indexHandler 和 helloHandler
	//根据不同的HTTP请求会调用不同的处理函数
	//访问 /，响应是URL.Path = /
	//而访问 /hello的响应则是请求头(header)中的键值对信息。
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	//第一个参数是地址，:9999表示在 9999 端口监听。而第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 输出 url 路径
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// 输出 header 信息
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
