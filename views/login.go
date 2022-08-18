package views

import (
	"net/http"
	"tb-go-blog/common"
	"tb-go-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
