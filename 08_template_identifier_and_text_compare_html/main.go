package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	t:=template.New("index.tmpl").
		Delims("{[","]}")
	//解析模板
	_,err:=t.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("Parse template error:",err)
		return
	}
	var name = "阿敦"
	//渲染模板
	err=t.Execute(w,name)
	if err != nil {
		fmt.Println("Render template error:",err)
		return
	}
}

func xss(w http.ResponseWriter,r *http.Request){
	//定义模板
	t:=template.New("xss.tmpl").
		Funcs(template.FuncMap{
			"safe":func(str string)template.HTML{
				return template.HTML(str)
			},
	})
	//解析模板
	_,err:=t.ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("Parse template error:",err)
		return
	}
	var str string= "<script>alert(123);</script>"
	str1:= "<a href='https://www.baidu.com/'>百度</a>"
	//渲染模板
	err=t.Execute(w,map[string]interface{}{
		"str":str,
		"str1":str1,
	})
	if err != nil {
		fmt.Println("Render template error:",err)
		return
	}
}

func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Println("Http Server start failed:",err)
		return
	}
}
