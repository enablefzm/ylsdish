package sovell

type ResOrderInfo struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
	List  []OrderInfo `json:"list"`
}

func NewResOrderInfo() ResOrderInfo {
	return ResOrderInfo{}
}
