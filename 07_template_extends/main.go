package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}

	var msg string = "阿敦"
	//渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板 注意:嵌套模板,被嵌套的后解析
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}
	var msg string = "阿敦"
	//渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}

	var msg string = "阿敦"
	//渲染模板
	err = t.ExecuteTemplate(w, "index2.tmpl", msg)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板 注意:父模板先解析
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}
	var msg string = "哈哈"
	//渲染模板
	err = t.ExecuteTemplate(w, "home2.tmpl", msg)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed :", err)
		return
	}
}
