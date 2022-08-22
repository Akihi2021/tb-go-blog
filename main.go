package main

import (
	"log"
	"net/http"
	"tb-go-blog/common"
	"tb-go-blog/router"
)

//用于测试json包装

//因为html缺函数实现，在这里临时实现一下
func init() {
	//模板加载
	common.LoadTemplate()
}
func main() {
	//程序入口
	//web程序 需要http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	//路由功能
	router.Router()
	//返回页面

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
