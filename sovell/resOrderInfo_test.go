package sovell

import (
	"fmt"
	"testing"
)

func TestGetOrders(t *testing.T) {

	req := NewReqOrders(GetTimeOnStr("2022-01-01 00:00:00"), GetTimeOnStr("2022-01-04 23:59:59"), 1)
	res, _ := GetOrders(req)
	fmt.Println(res)

	reqDetail := NewReqOrderDetail("637740417889057895")
	resDetail, _ := GetOrderDetail(reqDetail)
	fmt.Println(resDetail)
}
