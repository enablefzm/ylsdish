package main

import (
	"fmt"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	// t.Log(dbs.Cfg.User)
	nowHouse := time.Now().Hour()
	if nowHouse > 6 && nowHouse < 22 {
		fmt.Println(nowHouse, "开始同步")
	} else {
		fmt.Println(nowHouse, "夜间不同步")
	}
}
