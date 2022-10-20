package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func say_hello(w http.ResponseWriter, r *http.Request){
	b,err:=ioutil.ReadFile("./02_http_demo/index.html")
	if err!=nil{
		fmt.Println("ReadFile error:",err)
		return
	}
	_,_=fmt.Fprintln(w,string(b))
}

func main() {
	http.HandleFunc("/hello",say_hello)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		fmt.Printf("http serve failed error:%v\n",err)
		return
	}
}
