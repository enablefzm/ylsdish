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

// 工作处理
func WorkOrders() error {
	// 获取最后时间
	obDb := dbs.NewMyDB()
	fmt.Println(obDb.GetLastTime().String())
	return nil
}

// 获取智能餐盘的业务列表
func GetOrders(obReq *ReqOrders) (ResOrderInfo, error) {
	obRes := NewResOrderInfo()
	// 生成请求的URL
	url := dbs.Cfg.DishServer + dbs.Cfg.OrdersPath
	bts, err := PostUrl(url, obReq)
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
	bts, err := PostUrl(url, obReq)
	if err != nil {
		return obRes, err
	}
	json.Unmarshal(bts, &obRes)
	return obRes, nil
}

func PostUrlStr(url string, req IReq) (string, error) {
	bts, err := PostUrl(url, req)
	if err != nil {
		return "", err
	}
	return string(bts), err
}

func PostUrl(url string, req IReq) ([]byte, error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(req.GetPost()))
	if err != nil {
		log.Println("访问", url, "服务器发生错误:", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
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
