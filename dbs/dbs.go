package dbs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ObDB *gorm.DB

func init() {
	/*
		err := LinkDb()
		if err != nil {
			panic("连接服务器发生误:" + err.Error())
		}
	*/
}

func LinkDb() error {
	var err error
	// dsn := "billManage:BillManage&2021@tcp(134.175.246.140:3316)/billManage?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		Cfg.User,
		Cfg.Pass,
		Cfg.Address,
		Cfg.Port,
		Cfg.DbName)
	ObDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDb, err1 := ObDB.DB()
	if err1 == nil {
		sqlDb.SetMaxIdleConns(Cfg.MinConn)
		sqlDb.SetMaxOpenConns(Cfg.MaxConn)
	}
	return err
}
