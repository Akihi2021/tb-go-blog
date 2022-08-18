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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1  1参数 分类的id 要取出参数
	path := r.URL.Path                        //此时path为 /c/1 要取出1
	cIdStr := strings.TrimPrefix(path, "/c/") //去掉 /c/ 前缀
	cId, err := strconv.Atoi(cIdStr)          //将字符串转换为int型
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	//获得文章列表，他可能会有分页
	//分页
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取数据出错：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")

	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
