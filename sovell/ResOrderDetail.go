package sovell

func NewResOrderDetail() ResOrderDetail {
	return ResOrderDetail{}
}

type ResOrderDetail struct {
	Code    int                     `json:"code"`
	Msg     string                  `json:"msg"`
	Seq     string                  `json:"seq"`
	Order   ResOrderDetailOrderInfo `json:"order"`
	Details []ResOrderDetailInfo    `json:"details"`
}

type ResOrderDetailOrderInfo struct {
	Seq        string `json:"seq"`
	CreateDate string `json:"create_date"`
	Type       string `json:"type"`
	AmtDues    string `json:"amt_dues"`
	Amt        string `json:"amt"`
	Qty        string `json:"qty"`
	Part       string `json:"part"`
	CardNo     string `json:"cardno"`
	Cid        string `json:"cid"`
	ShopName   string `json:"shopname"`
	TermName   string `json:"termname"`
}

type ResOrderDetailInfo struct {
	ItemType string `json:"type"`
	Pid      string `json:"pid"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Amt      string `json:"amt"`
	Weight   string `json:"weight"`
}
