package router

import (
	"net/http"
	"tb-go-blog/api"
	"tb-go-blog/views"
)

func Router() {
	//1.返回页面 views 2.返回数据（json） 3.静态资源
	http.HandleFunc("/", views.HTML.Index) //需要两个参数，用index函数来给 用来相应根目录
	//http://localhost:8080/c/1  1参数 分类的id 要取出参数
	http.HandleFunc("/c/", views.HTML.Category)
	//登录
	http.HandleFunc("/login", views.HTML.Login)
	//文章详情
	//http://localhost:8080/p/7
	http.HandleFunc("/p/", views.HTML.Detail)
	//写文章页面
	http.HandleFunc("/writing", views.HTML.Writing)
	//登录接口实现
	http.HandleFunc("/api/v1/login", api.API.Login)
	//发布和编辑文章
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	//归档
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	//让/resource/路径请求的东西映射到public/resource/
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
