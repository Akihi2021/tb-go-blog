package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

//配置
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string //导航条
	Bilibili    string   //bilibili链接地址
	Avatar      string   //头像
	UserName    string
	UserDesc    string //用户描述
} //不放数据库，放在配置文件config.toml里

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

//配置文件的启动
var Cfg *tomlConfig

func init() {
	//程序启动的时候就会执行init方法
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "tb-go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir //当前目录
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err) //直接抛出，让他死掉
	}
}
