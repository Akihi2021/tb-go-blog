package api

import (
	"net/http"
	"tb-go-blog/common"
	"tb-go-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	//接收用户名 密码 返回对应的json数据
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passWord := params["passwd"].(string)
	loginRes, err := service.Login(userName, passWord)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
