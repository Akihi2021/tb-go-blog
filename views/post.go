package views

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"tb-go-blog/common"
	"tb-go-blog/service"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	//获取路径参数
	path := r.URL.Path                        //此时path为 /p/7.html 要取出1
	pIdStr := strings.TrimPrefix(path, "/p/") //去掉 /p/ 前缀
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html") //去掉.html后缀
	//7
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	postRes, err := service.GetPostDetail(pid)
	if err != nil {
		detail.WriteError(w, errors.New("查询出错"))
		return
	}
	detail.WriteData(w, postRes)
}
