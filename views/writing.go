package views

import (
	"net/http"
	"tb-go-blog/common"
	"tb-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := common.Template.Writing
	wr := service.Writing()
	writing.WriteData(w, wr)
}
