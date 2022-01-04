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
	DishKey         string
	DishShop        int
	DishTerm        int
	DishOper        string
	DishPageSize    int
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
		DishServer:      "http://27.154.237.130:12510",
		OrdersPath:      "/xdf_jhnoa/orders",
		OrderDetailPath: "/xdf_jhnoa/orders",
		DishKey:         "b699e89068304252b0fcea2a001b8eff",
		DishShop:        101,
		DishTerm:        103,
		DishOper:        "yls",
		DishPageSize:    40,
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
				case "dish_key":
					cfg.DishKey = v
				case "dish_shop":
					cfg.DishShop = vatools.SInt(v)
				case "dish_term":
					cfg.DishTerm = vatools.SInt(v)
				case "dish_oper":
					cfg.DishOper = v
				case "dish_page_size":
					cfg.DishPageSize = vatools.SInt(v)
					if cfg.DishPageSize < 1 {
						cfg.DishPageSize = 1
					}
					if cfg.DishPageSize > 100 {
						cfg.DishPageSize = 100
					}
				}
			}
		}
	}
	return cfg
}
