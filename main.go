package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tb-go-blog/config"
	"tb-go-blog/models"
	"time"
)

//用于测试json包装
type IndexData struct {
	Title string `json:"title"` //这样返回的json就不是Title而是title了，符合规范
	Desc  string
}

//因为html缺函数实现，在这里临时实现一下
func IsODD(num int) bool {
	//用于判断是否是偶数，用在header.html 13行处
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1] //用于返回对应路径
	//因为数组是这么存的["首页","/", "GO语言","/golang", "归档","/pigeonhole", "关于","/about"]
	//+1就能返回对应的路径
}

//返回当前时间
func Date(layout string) string {
	return time.Now().Format(layout)
}
func index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json") //用于将返回的信息用json包装，这里包装了的话会把整个页面都变成json代码返回
	var indexData IndexData
	indexData.Title = "唐博的第一个go博客"
	indexData.Desc = "用于测试json包装"

	t := template.New("index.html")
	//拿到当前路径
	//path, _ := os.Getwd()
	path := config.Cfg.System.CurrentDir
	//解析路径，访问博客首页时因为有多个页面的嵌套，需要将其涉及的所有模板都进行解析
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	//header.html要用到这个函数判断偶数
	//header.html要用到这个函数获得路径
	//footer.html要用到这个函数获得时间
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错：", err)
	}
	//页面上所有的数据必须有定义
	//来点假数据，要写的太多了，选择使用配置文件生成
	/*	viewer := config.Viewer{
		Title:       "TB的博客",
		Description: "TB的第一个GO博客！！",
		Logo:        "",
		Navigation:  nil,
		Bilibili:    "",
		Avatar:      "",
		UserName:    "",
		UserDesc:    "",
	}*/
	//首页的假数据
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "巴拉巴拉",
			UserName:     "TB",
			ViewCount:    123,
			CreateAt:     "2022-08-16",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	//执行，打印错误
	err = t.Execute(w, hr)
	fmt.Println(err)
}

func main() {
	//程序入口
	//web程序 需要http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}
	//路由功能
	http.HandleFunc("/", index) //需要两个参数，用index函数来给 用来相应根目录
	//让/resource/路径请求的东西映射到public/resource/
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	//返回页面

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
