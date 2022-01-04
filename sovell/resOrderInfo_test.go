package sovell

import (
	"fmt"
	"testing"
)

func TestGetOrders(t *testing.T) {
	// res, _ := GetOrders()
	// t.Log(res)
	//req := NewReqOrders(GetTimeOnStr("2021-12-01 00:00:00"), GetTimeOnStr("2022-01-04 23:59:59"), 1)
	//res, _ := GetOrders(req)
	//fmt.Println(res)

	req := NewReqOrderDetail("637740417889057895")
	res, _ := GetOrderDetail(req)
	fmt.Println(res)
}
