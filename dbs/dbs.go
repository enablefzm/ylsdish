package dbs

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ObDB *gorm.DB

func LinkDb() error {
	var err error
	// dsn := "billManage:BillManage&2021@tcp(134.175.246.140:3316)/billManage?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		Cfg.User,
		Cfg.Pass,
		Cfg.Address,
		Cfg.Port,
		Cfg.DbName)
	ObDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

// 获取最后同步的时间
func GetLastTime() time.Time {
	// 读取本地保存的数据
	// t, err := time.ParseInLocation(vatools.TIME_FORMAT, )
	mydb := NewMyDB()
	return mydb.GetLastTime()
}
