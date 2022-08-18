package models

//从数据库中取
//home.html里有个人信息，在个人信息personal.html里要用到这个，分类循环
type Category struct {
	Cid      int
	Name     string
	CreateAt string //创建时间
	UpdateAt string //更新时间
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
