package sovell

import (
	"fmt"
	"github.com/enablefzm/gotools/vatools"
	"strings"
	"ylsdish/dbs"
)

type ReqOrderDetail struct {
	v      int    // 版本号
	d      string // 发起时间
	format string // 返回的数据格式
	shop   int    // 消费ID
	term   int    // 终端编号
	oper   string // 终端操作员
	sig    string // 签名
	seq    string // 要查询的业务流水号
}

func NewReqOrderDetail(req string) *ReqOrderDetail {
	return &ReqOrderDetail{
		v:      3,
		d:      vatools.GetNowTimeString(),
		format: "json",
		shop:   dbs.Cfg.DishShop,
		term:   dbs.Cfg.DishTerm,
		oper:   dbs.Cfg.DishOper,
		sig:    "",
		seq:    req,
	}
}

func (req *ReqOrderDetail) GetPost() string {
	arr := []string{
		"t=21",
		fmt.Sprintf("v=%d", req.v),
		fmt.Sprintf("d=%s", req.d),
		fmt.Sprintf("format=%s", req.format),
		fmt.Sprintf("shop=%d", req.shop),
		fmt.Sprintf("term=%d", req.term),
		fmt.Sprintf("oper=%s", req.oper),
		fmt.Sprintf("seq=%s", req.seq),
	}
	req.sig = GetSign(arr)
	arrPost := append(arr, "sig="+req.sig)
	return strings.Join(arrPost, "&")
}
