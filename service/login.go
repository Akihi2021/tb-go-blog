package service

import (
	"errors"
	"tb-go-blog/dao"
	"tb-go-blog/models"
	"tb-go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginResponse, error) {
	//md5加密
	passwd = utils.Md5Crypt(passwd, "mszlu")
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uId := user.Uid
	//生成token 使用jwt技术进行生成 令牌 A.B.C
	token, err := utils.Award(&uId)
	if err != nil {
		return nil, errors.New("token未生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginResponse{
		token,
		userInfo,
	}
	return lr, nil
}
