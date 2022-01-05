package sovell

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"ylsdish/dbs"

	"github.com/enablefzm/gotools/vatools"
)

// 获取智能餐盘的业务列表
func GetOrders(obReq *ReqOrders) (ResOrderInfo, error) {
	obRes := NewResOrderInfo()
	// 生成请求的URL
	url := dbs.Cfg.DishServer + dbs.Cfg.OrdersPath
	bts, err := TryPostUrl(url, obReq)
	if err != nil {
		return obRes, err
	}
	json.Unmarshal(bts, &obRes)
	return obRes, nil
}

// 获取智能餐盘流水号详细数据信息
func GetOrderDetail(obReq *ReqOrderDetail) (ResOrderDetail, error) {
	obRes := NewResOrderDetail()
	// 生成请求的URL
	url := dbs.Cfg.DishServer + dbs.Cfg.OrderDetailPath
	bts, err := TryPostUrl(url, obReq)
	if err != nil {
		return obRes, err
	}
	json.Unmarshal(bts, &obRes)
	return obRes, nil
}

// 可以尝试3次错误的请求
func TryPostUrl(url string, req IReq) ([]byte, error) {
	var btVal []byte
	var err error
	for i := 0; i < 3; i++ {
		btVal, err = PostUrl(url, req)
		if err == nil {
			break
		}
		log.Println("访问服务器", url, "发生错误:", err.Error())
		if i < 3 {
			log.Println("等待5秒后重试...")
			time.Sleep(time.Second * 5)
		}
	}
	return btVal, err
}

func PostUrl(url string, req IReq) ([]byte, error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(req.GetPost()))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 获取签名
func GetSign(kv []string) string {
	sort.Strings(kv)
	kv = append(kv, dbs.Cfg.DishKey)
	strJoin := strings.Join(kv, "")
	return vatools.MD5(strJoin)
}

// 通过字符串获取时间对象
func GetTimeOnStr(strTime string) time.Time {
	t, err := time.ParseInLocation(vatools.TIME_FORMAT, strTime, time.Local)
	if err != nil {
		fmt.Println(err.Error())
		return time.Now()
	}
	return t
}
