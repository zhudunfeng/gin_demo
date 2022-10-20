package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age int
	Gender string
}

func sayHello(w http.ResponseWriter,r *http.Request){
	//1.定义模板
	//2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("Parse template error:",err)
		return
	}
	//3.渲染模板
	//定义填充数据
	u1:=User{
		Name: "阿敦",
		Age :18,
		Gender: "男",
	}

	m1:=map[string]interface{}{
		"name":"莉莉",
		"age":18,
		"gender":"女",
	}

	hobbyList:=[]string{
		"唱",
		"跳",
		"rapper",
		"篮球",
	}

	//err=t.Execute(w,u1)
	err=t.Execute(w, map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})

	if err!=nil{
		fmt.Println("Render template error:",err)
		return
	}
}

func main() {
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("HTTP Server started failed:",err)
		return
	}
}
