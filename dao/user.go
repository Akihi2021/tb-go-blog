package dao

import (
	"log"
	"tb-go-blog/models"
)

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	if row.Err() != nil {
		log.Println(row.Err())
		return ""
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}
func GetUser(userName, passwd string) *models.User {
	row := DB.QueryRow(
		"select * from blog_user where user_name = ? and passwd = ? limit 1", userName, passwd)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var user = &models.User{}
	//把相关信息扫描到user里
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		return nil
	}
	return user
}
