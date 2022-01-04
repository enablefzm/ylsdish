package sovell

// 智能餐盘的订单列表信息
type OrderInfo struct {
	Seq        string  `json:"seq"`
	Type       int     `json:"type"`
	CreateDate string  `json:"create_date"`
	Qty        int     `json:"qty"`
	Amt        float64 `json:"amt"`
	Part       int     `json:"part"`
}
