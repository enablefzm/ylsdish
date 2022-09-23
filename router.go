package main

import (
	"fmt"
	"log"
	"net/http"
	"ylsdish/dbs"
	"ylsdish/yls"

	"github.com/enablefzm/gotools/vatools"
)

// 初始化路由信息
func InitRouter() {
	// 建立Http服务
	http.HandleFunc("/startsync", func(wr http.ResponseWriter, req *http.Request) {
		go yls.WorkSyncDish()
		wr.Write([]byte(vatools.GetNowTimeString() + " ====> 手动执行了同步操作！！"))
	})
	// 设定同步时间
	http.HandleFunc("/setwaittime", func(wr http.ResponseWriter, req *http.Request) {
		arg := req.URL.Query().Get("t")
		waitTime := vatools.SInt(arg)
		if waitTime < 1 {
			waitTime = 1
		}
		dbs.Cfg.DishWaitTime = waitTime
		log.Println("设定同步等待时间为：", dbs.Cfg.DishWaitTime, "分钟")
		wr.Write([]byte(fmt.Sprint("设定当前同步等待时间为:", waitTime)))
	})
}
