package ylsdb

// 美食数据
type Appdishe struct {
	Id       string
	Name     string
	Code     string
	ZpDishId int `gorm:"column:ZpDishId"`
}
