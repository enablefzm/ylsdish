package main

import (
	"log"
	"ylsdish/dbs"
)

const ver string = "1.0.0"

func main() {
	log.Println("运行象屿慈爱养老信息管理系统【餐厅消费数据同步服务】Ver", ver)
	log.Println("开始连接数据库...")
	err := dbs.LinkDb()
	if err != nil {
		log.Println("连接数据库失败")
		log.Println(err.Error())
		return
	}
	log.Println("连接成功!")
}
