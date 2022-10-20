package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//测试模板自定义函数
func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数 kua
	// 要么只有一个返回值,要么有两个返回值,第二个返回值必须是error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}

	//定义模板
	t := template.New("f.tmpl") //创建一个名字是f的模板对象
	//告诉模板引擎,我现在多了一个自定义的函数k
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	//解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}

	name := "阿敦"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

//测试嵌套模板
func f2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板 注意:嵌套模板,被嵌套的后解析
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("Parse template error:", err)
		return
	}
	var name string = "阿敦"
	//渲染模板
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("Render template error:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpldemo", f2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP Server start failed :", err)
		return
	}
}
