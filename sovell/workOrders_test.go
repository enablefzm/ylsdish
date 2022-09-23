package sovell

import (
	"fmt"
	"testing"
	"time"
	"ylsdish/dbs"
)

func TestWorkOrders(t *testing.T) {
	/*
		req := NewReqProd(110)
		btVal, err := PostUrl(dbs.Cfg.DishServer+"/xdf_jhnoa/prod", req)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(btVal))
		}
	*/
	nowTime := time.Now()
	nextTime := nowTime.Add(-1 * time.Second)
	fmt.Println(nowTime, nextTime)

	mydb := dbs.NewMyDB()
	fmt.Println(mydb.GetLastTime())
	mydb.SetLastTimeSave(nowTime)
}
