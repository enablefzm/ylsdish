package dbs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ObDB *gorm.DB

func LinkDb() error {
	var err error
	dsn := "billManage:BillManage&2021@tcp(134.175.246.140:3316)/billManage?charset=utf8mb4&parseTime=True"
	ObDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
