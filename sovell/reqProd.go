package sovell

import (
	"fmt"
	"github.com/enablefzm/gotools/vatools"
	"strings"
	"ylsdish/dbs"
)

type ReqProd struct {
	v      int
	d      string
	format string
	shop   int
	term   int
	oper   string
	id     int
	sig    string
}

func NewReqProd(id int) *ReqProd {
	return &ReqProd{
		v:      3,
		d:      vatools.GetNowTimeString(),
		format: "json",
		shop:   dbs.Cfg.DishShop,
		term:   dbs.Cfg.DishTerm,
		oper:   dbs.Cfg.DishOper,
		id:     id,
		sig:    "",
	}
}

func (req *ReqProd) GetPost() string {
	arr := []string{
		fmt.Sprintf("v=%d", req.v),
		fmt.Sprintf("d=%s", req.d),
		fmt.Sprintf("format=%s", req.format),
		fmt.Sprintf("shop=%d", req.shop),
		fmt.Sprintf("term=%d", req.term),
		fmt.Sprintf("oper=%s", req.oper),
		fmt.Sprintf("id=%d", req.id),
	}
	req.sig = GetSign(arr)
	arrPost := append(arr, "sig="+req.sig)
	return strings.Join(arrPost, "&")
}
