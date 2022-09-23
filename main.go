package main

import (
	"fmt"
	"github.com/enablefzm/gotools/vatools"
	"log"
	"net/http"
	"time"
	"ylsdish/dbs"
	"ylsdish/yls"
)

// v1.0.1
//	修改增加可以动态设置同步时间
// 	只有在工作时段内早上6点至晚上23点内才会同步数据

const ver string = "1.0.1"

func main() {
	log.Println("运行象屿慈爱养老信息管理系统【餐厅消费数据同步服务】Ver", ver)
	log.Println("开始连接数据库...")
	err := dbs.LinkDb()
	if err != nil {
		log.Println("连接数据库失败：", err.Error())
		return
	}
	log.Println("连接成功!")
	// 启动服务
	go func() {
		time.Sleep(2 * time.Second)
		for {

			// 开始同步数据
			// yls.WorkSyncDish()
			// 只在工作时间内同步
			yls.WorkSyncDishOnWorkTime()
			log.Println("开始等待", dbs.Cfg.DishWaitTime, "分钟后再继续同步")
			// 等待时间
			time.Sleep(time.Minute * time.Duration(dbs.Cfg.DishWaitTime))
		}
	}()
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
	log.Println("运行HttpServer服务")
	err = http.ListenAndServe(":"+dbs.Cfg.HttpServerPort, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
