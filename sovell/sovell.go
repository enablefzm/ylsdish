package sovell

import (
	"github.com/enablefzm/gotools/vatools"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const key string = "aabbccddeeffgg"

// 获取智能餐盘的业务列表
func GetOrders() (ResOrderInfo, error) {
	obRes := NewResOrderInfo()
	// 生成请求的URL
	url := "http://xxxx:19001/xdf_jhnoa/orders"
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("xxx=aaa"))
	if err != nil {
		log.Println("请求订单列表发生错误:", err.Error())
		return obRes, err
	}
	defer res.Body.Close()
	// 读取内容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return obRes, nil
	}
	log.Println(string(body))
	return obRes, nil
}

// 获取签名
func GetSign(kv []string) string {
	kv = append(kv, key)
	strJoin := strings.Join(kv, "")
	return vatools.MD5(strJoin)
}
