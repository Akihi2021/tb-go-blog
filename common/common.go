package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"tb-go-blog/config"
	"tb-go-blog/models"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	//优化操作
	//声明一个 WaitGroup，通过 Add 方法设置一个计数器的值，需要跟踪多少协程就设置多少。
	//每个协程在执行完毕的时候，一定要调 Done 方法，让计数器减1，告诉 WaitGroup 该协程已经执行完毕。
	//最后调用 Wait 方法，一直等待，直到计数器的值变为0，也就是所有跟踪的协程执行完毕了。
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}
