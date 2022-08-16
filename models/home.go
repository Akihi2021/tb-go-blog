package models

import "tb-go-blog/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category //定义分类
	Posts     []PostMore //文章相关信息
	Total     int        //pagination.html需要的分页的总数
	Page      int        //pagination.html需要的每页的数量
	Pages     []int
	PageEnd   bool //pagination.html需要的判断是不是结束了
}
