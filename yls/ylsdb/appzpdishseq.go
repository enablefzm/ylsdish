package ylsdb

import "time"

type Appzpdishseq struct {
	Id         string
	Seq        string
	CreateDate time.Time
	Type       string
	Amt        string
	Qty        string
	Part       string
	State      int
}
