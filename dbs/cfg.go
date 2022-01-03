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
		User:            "root",
		Pass:            "Health2021*",
		Address:         "123.139.116.237",
		Port:            "3306",
		DbName:          "platform_pub",
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
				case "port":
					cfg.Port = v
				case "dbname":
					cfg.DbName = v
				case "maxconn":
					maxConn := vatools.SInt(v)
					if maxConn < 1 {
						maxConn = 1
					}
					cfg.MaxConn = vatools.SInt(v)
				case "minconn":
					minConn := vatools.SInt(v)
					if minConn < 1 {
						minConn = 1
					}
					cfg.MinConn = minConn
				}
			}
		}
		if mp, ok := c.GetNode("dish"); ok {
			for k, v := range mp {
				switch k {
				case "server":
					cfg.DishServer = v
				case "order_path":
					cfg.OrdersPath = v
				case "order_detail_path":
					cfg.OrderDetailPath = v
				}
			}
		}
	}
	return cfg
}
