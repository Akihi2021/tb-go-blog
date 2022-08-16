package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//用于测试json包装
type IndexData struct {
	Title string `json:"title"` //这样返回的json就不是Title而是title了，符合规范
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json") //用于将返回的信息用json包装，这里包装了的话会把整个页面都变成json代码返回
	var indexData IndexData
	indexData.Title = "唐博的第一个go博客"
	indexData.Desc = "用于测试json包装"

	t := template.New("index.html")
	//拿到当前路径
	path, _ := os.Getwd()
	//解析路径
	t, _ = t.ParseFiles(path + "/template/index.html")
	//执行，打印错误
	err := t.Execute(w, indexData)
	fmt.Println(err)
}

func main() {
	//程序入口
	//web程序 需要http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}

	http.HandleFunc("/", index) //需要两个参数，用index函数来给 用来相应根目录
	//返回页面

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
