package ylsdb

import "time"

type Appdishconfirm struct {
	Id             string
	MemberId       string    `gorm:"column:MemberId"`       // 长者ID
	DishId         string    `gorm:"column:DishId"`         // 食物对应的ID
	ItemType       int       `gorm:"column:ItemType"`       // 餐食时段
	OperateId      string    `gorm:"column:OperateId"`      // 操作员ID
	ConfirmOperate string    `gorm:"column:ConfirmOperate"` // 食物加工类型
	CreationTime   time.Time `gorm:"column:CreationTime"`   // 创建时间
	ConfirmDate    time.Time `gorm:"column:ConfirmDate"`    // 用餐时间
	Count          int       `gorm:column:Count`            // 数量
	DishName       string    `gorm:"column:DishName"`       // 食物名称
}
