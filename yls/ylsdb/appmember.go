package ylsdb

type Appmember struct {
	Id            string
	FullName      string `gorm:"column:FullName"`
	IcCardAccount string `gorm:"column:IcCardAccount"`
}
