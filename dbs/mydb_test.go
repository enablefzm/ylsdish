package dbs

import (
	"fmt"
	"testing"
)

func TestMydb(t *testing.T) {
	ob := NewMyDB()
	fmt.Println(ob.GetLastTime())
	ob.Save()
}
