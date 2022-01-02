package sovell

import "time"

type ReqOrders struct {
	v      int       // 请求版本号
	d      time.Time // 发起请求时间
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
