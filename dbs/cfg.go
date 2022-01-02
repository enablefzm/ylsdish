package dbs

import (
	"github.com/enablefzm/gotools/vaini"
	"github.com/enablefzm/gotools/vatools"
)

var Cfg *YlsDishCfg = NewYlsDishCfg()

type YlsDishCfg struct {
	User            string
	Pass            string
	Address         string
	Port            string
	DbName          string
	MaxConn         int
	MinConn         int
	DishServer      string
	OrdersPath      string
	OrderDetailPath string
}

func NewYlsDishCfg() *YlsDishCfg {
	cfg := &YlsDishCfg{
		User:            "billManage",
		Pass:            "BillManage&2021",
		Address:         "134.175.246.140",
		Port:            "3316",
		DbName:          "billManage",
		MaxConn:         10,
		MinConn:         1,
		DishServer:      "http://127.0.0.1:19001",
		OrdersPath:      "/xdf_jhnoa/orders",
		OrderDetailPath: "/xdf_jhnoa/orders",
	}
	// 加载文件并解析
	path, err := vatools.GetNowPath()
	if err == nil {
		c := vaini.NewConfig(path + "/cfg.ini")
		if mp, ok := c.GetNode("database"); ok {
			for k, v := range mp {
				switch k {
				case "user":
					cfg.User = v
				case "pass":
					cfg.Pass = v
				case "address":
					cfg.Address = v
				}
			}
		}
	}
	return cfg
}
