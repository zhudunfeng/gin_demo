package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//1、定义模板
	//2、解析模板
	t, err := template.ParseFiles("./04_http_template/hello.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}
	//3、渲染模板
	name := "阿敦"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("render template error:", err)
		return
	}
}

func main() {
	//映射处理器
	http.HandleFunc("/", sayHello)
	//监听端口
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed error:", err)
		return
	}
}
