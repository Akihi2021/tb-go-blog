package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

//用于处理模板

type TemplateBlog struct {
	*template.Template
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog //归档页面
	Writing    TemplateBlog
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (t *TemplateBlog) WriteError(w io.Writer, err error) {
	//err := t.Execute(w, data)
	if err != nil {
		_, err := w.Write([]byte(err.Error()))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
func InitTemplate(templateDir string) (HtmlTemplate, error) {

	tp, err := readTemplate([]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)
	var htmlTemplate HtmlTemplate
	if err != nil {
		return htmlTemplate, err
	}
	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func IsODD(num int) bool {
	//用于判断是否是偶数，用在header.html 13行处
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1] //用于返回对应路径
	//因为数组是这么存的["首页","/", "GO语言","/golang", "归档","/pigeonhole", "关于","/about"]
	//+1就能返回对应的路径
}

func Date(layout string) string {
	//返回当前时间
	return time.Now().Format(layout)
}

func DateDay(date time.Time) string {
	//格式化时间，go的诞生时间
	return date.Format("2006-01-02 15:04:05")
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var tbs []TemplateBlog
	for _, view := range templates {
		viewname := view + ".html"
		t := template.New(viewname)
		//拿到当前路径
		//path, _ := os.Getwd()
		//path := config.Cfg.System.CurrentDir
		//解析路径，访问博客首页时因为有多个页面的嵌套，需要将其涉及的所有模板都进行解析
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		post := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		//header.html要用到这个函数判断偶数
		//header.html要用到这个函数获得路径
		//footer.html要用到这个函数获得时间
		//custom.html要用到这个函数格式化时间
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewname, home, header, footer, personal, post, pagination)
		if err != nil {
			log.Println("解析模板出错：", err)
			return nil, err
		}
		//执行，打印错误,不需要了，只需要返回模板
		//err = t.Execute(w, hr)
		//fmt.Println(err)
		var tb TemplateBlog
		tb.Template = t
		tbs = append(tbs, tb)
	}
	return tbs, nil

}
