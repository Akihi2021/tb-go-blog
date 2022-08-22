package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"tb-go-blog/common"
	"tb-go-blog/service"
)

type IndexData struct {
	Title string `json:"title"` //这样返回的json就不是Title而是title了，符合规范
	Desc  string
}

func index() {

}
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index

	//w.Header().Set("Content-Type", "application/json") //用于将返回的信息用json包装，这里包装了的话会把整个页面都变成json代码返回
	var indexData IndexData
	indexData.Title = "唐博的第一个go博客"
	indexData.Desc = "用于测试json包装"

	//页面上所有的数据必须有定义
	//数据库查询

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
	//分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	index.WriteData(w, hr)

}
