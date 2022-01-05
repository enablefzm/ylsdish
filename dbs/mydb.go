package dbs

import (
	"time"

	"github.com/enablefzm/gotools/vatools"
)

const FILE_MYDB_PATH string = "./mydb.db"

type MyDB struct {
	LastTime string `json:"lasttime"`
}

func NewMyDB() *MyDB {
	mydb := MyDB{}
	arrStr, _ := vatools.ReadFileLR(FILE_MYDB_PATH)
	if len(arrStr) > 0 {
		strJson := arrStr[0]
		vatools.UnJson(strJson, &mydb)
	}
	return &mydb
}

func (mydb *MyDB) GetLastTime() time.Time {
	if len(mydb.LastTime) < 8 {
		mydb.LastTime = "2021-10-10 00:00:00"
	}
	t, _ := time.ParseInLocation(vatools.TIME_FORMAT, mydb.LastTime, time.Local)
	return t
}

func (mydb *MyDB) SetLastTime(dt time.Time) {
	mydb.LastTime = dt.Format(vatools.TIME_FORMAT)
}

func (mydb *MyDB) Save() error {
	strJson, err := vatools.Json(mydb)
	vatools.WriteFileLR(FILE_MYDB_PATH, strJson)
	return err
}

func (mydb *MyDB) SetLastTimeSave(dt time.Time) {
	mydb.SetLastTime(dt)
	mydb.Save()
}
