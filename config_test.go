package main

import (
	"testing"
	"ylsdish/dbs"
)

func TestConfig(t *testing.T) {
	t.Log(dbs.Cfg.DbName)
}
