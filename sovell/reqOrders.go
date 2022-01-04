package sovell

import (
	"fmt"
	"github.com/enablefzm/gotools/vatools"
	"strings"
	"time"
	"ylsdish/dbs"
)

type ReqOrders struct {
	v      int       // 请求版本号
	d      string    // 发起请求时间
	format string    // 返回的数据格式 默认都为json
	shop   int       // 餐厅编号
	term   int       // 终端编号
	oper   string    // 终端操作员
	sig    string    // 请求签名
	t      int       // 请求类型 11 业务列表
	pi     int       // 当前要取的业码
	ps     int       // 当前取值的业大小
	sd     time.Time // 业务开始时间
	ed     time.Time // 业务结束时间
}

func NewReqOrders(startDt, endDt time.Time, nowPage int) *ReqOrders {
	if nowPage < 1 {
		nowPage = 1
	}
	if endDt.Before(startDt) {
		endDt = startDt.Add(time.Second * 1)
	}
	return &ReqOrders{
		v:      3,
		d:      vatools.GetNowTimeString(),
		format: "json",
		shop:   dbs.Cfg.DishShop,
		term:   dbs.Cfg.DishTerm,
		oper:   dbs.Cfg.DishOper,
		sig:    "",
		t:      11,
		pi:     nowPage,
		ps:     dbs.Cfg.DishPageSize,
		sd:     startDt,
		ed:     endDt,
	}
}

func (req *ReqOrders) GetPost() string {
	arr := []string{
		fmt.Sprintf("v=%d", req.v),
		fmt.Sprintf("d=%s", req.d),
		fmt.Sprintf("format=%s", req.format),
		fmt.Sprintf("shop=%d", req.shop),
		fmt.Sprintf("term=%d", req.term),
		fmt.Sprintf("oper=%s", req.oper),
		fmt.Sprintf("t=%d", req.t),
		fmt.Sprintf("pi=%d", req.pi),
		fmt.Sprintf("ps=%d", req.ps),
		fmt.Sprintf("sd=%s", req.sd.Format(vatools.TIME_FORMAT)),
		fmt.Sprintf("ed=%s", req.ed.Format(vatools.TIME_FORMAT)),
	}
	// 计算KEY
	sig := GetSign(arr)
	arrJson := append(arr, "sig="+sig)
	return strings.Join(arrJson, "&")
}
