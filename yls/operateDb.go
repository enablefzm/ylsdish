package yls

import (
	"fmt"
	"github.com/enablefzm/gotools/guid"
	"github.com/enablefzm/gotools/vatools"
	"ylsdish/dbs"
	"ylsdish/sovell"
	"ylsdish/yls/ylsdb"
)

// 通过长者的饭卡ID来获取长者数据
func DbGetMember(cid string) (*ylsdb.Appmember, error) {
	var rss []ylsdb.Appmember
	result := dbs.ObDB.Where("IcCardAccount = ?", cid).Limit(2).Find(&rss)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(rss) < 1 {
		return nil, fmt.Errorf("JNULL")
	}
	return &rss[0], nil
}

// 写入智盘长者消费的数据
func DbWriteOrderInfo(orderInfo *sovell.ResOrderDetailOrderInfo, state EmSeqState) error {
	v := &ylsdb.Appzpdishseq{
		Id:         guid.NewString(),
		Seq:        orderInfo.Seq,
		CreateDate: vatools.STime(orderInfo.CreateDate),
		Type:       orderInfo.Type,
		Amt:        orderInfo.Amt,
		Qty:        orderInfo.Qty,
		Part:       orderInfo.Part,
		State:      int(state),
	}
	result := dbs.ObDB.Create(v)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 通过智盘的食物ID来获取慈爱系统的食物信息
func DbGetDishInfo(zpDishID int) (*ylsdb.Appdishe, error) {
	var rss []ylsdb.Appdishe
	result := dbs.ObDB.Where("ZpDishId = ?", zpDishID).Limit(2).Find(&rss)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(rss) < 1 {
		return nil, fmt.Errorf("JNULL")
	}
	return &rss[0], nil
}

// 写入智盘长者消费的餐食数据
// func DbWriteDishConfirm(pMember *ylsdb.Appmember, info *sovell.ResOrderDetailOrderInfo)
