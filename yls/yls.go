package yls

import (
	"github.com/enablefzm/gotools/guid"
	"github.com/enablefzm/gotools/vatools"
	"gorm.io/gorm"
	"log"
	"time"
	"ylsdish/dbs"
	"ylsdish/sovell"
	"ylsdish/yls/ylsdb"
)

var ICountSeq int = 0       // 总共导入单据数量
var ICountSeqDetail int = 0 // 总共导入单据明细数量
var ILoopCountSeq int = 0
var ILoopCountSeqDetail int = 0

// 执行同步数据
func WorkSyncDish() {
	ILoopCountSeq = 0
	ILoopCountSeqDetail = 0
	// 获取最后一次同步的时间
	mydb := dbs.NewMyDB()
	lastTime := mydb.GetLastTime()
	lastTime = vatools.STime("2022-01-04 00:00:00")
	nowTime := time.Now()
	// 获取服务端数据
	var iPage int = 1
	var iStep int = 0
	log.Println("Start -> 开始同步餐盘数据", lastTime.Format(vatools.TIME_FORMAT), " 至 ", nowTime.Format(vatools.TIME_FORMAT), "...")
	// 获取数据
	for iStep < 1000 {
		iStep++
		log.Println("准备获取第 [", iPage, "] 页数据...")
		obReq := sovell.NewReqOrders(lastTime, nowTime, iPage)
		obRes, err := sovell.GetOrders(obReq)
		if err != nil {
			log.Println("访问", dbs.Cfg.DishServer, "服务器发生错误:", err.Error())
			log.Println("退出数据同步")
			break
		}
		// 执行结果
		nowIdx := (iPage - 1) * dbs.Cfg.DishPageSize
		nowSize := len(obRes.List)
		log.Println("获取成功！总共获取[", obRes.Total, "]条需要同步的业务数量，准备同步第", nowIdx, "-", nowIdx+nowSize, "...")
		for i := 0; i < nowSize; i++ {
			time.Sleep(time.Millisecond * 500)
			WorkSyncDishDetail(obRes.List[i].Seq)
		}
		time.Sleep(time.Second * 1)
		if nowSize < dbs.Cfg.DishPageSize {
			// 同步完成退出
			log.Println("Success -> 同步数据完成！")
			mydb.SetLastTimeSave(nowTime)
			break
		}
		iPage++
	}
	log.Println("========================== Report Detail =============================")
	log.Println("本次导入【", ILoopCountSeq, "】条业务单，长者消费明细数量为【", ILoopCountSeqDetail, "】条")
	log.Println("自服务启动合计导入业务量为【", ICountSeq, "】条，长者消费明细为【", ICountSeqDetail, "】条")
}

func WorkSyncDishDetail(seq string) {
	log.Println("->开始同步业务号", seq, "...")
	obReq := sovell.NewReqOrderDetail(seq)
	obRes, err := sovell.GetOrderDetail(obReq)
	if err != nil {
		log.Println("访问", dbs.Cfg.DishServer, " 请求业务单号:", seq, "发生错误:", err.Error())
		return
	}
	log.Println("请求服务器返回", len(obRes.Details), "条详细数据，开始导入数据...")
	// 判断是否有同步过这个业务单号
	var rss []ylsdb.Appzpdishseq
	dbs.ObDB.Where("seq = ?", seq).Limit(2).Find(&rss)
	if len(rss) > 0 {
		log.Println("业务号", seq, "已经存在！跳过！")
		return
	}
	// 获取长者信息数据
	pMember, err := DbGetMember(obRes.Order.Cid)
	if err != nil {
		if err.Error() == "JNULL" {
			// 长者不存在，标记这个业务单为长者不存
			DbWriteOrderInfo(&obRes.Order, SEQ_NOMEMBER)
		}
		log.Println("这个业务单号:", seq, " 有错误:", err.Error())
		return
	}
	// TODO 执行写入慈爱管理系统
	var seqState EmSeqState = SEQ_SUCCESS
	// 创建事务
	err = dbs.ObDB.Transaction(func(tx *gorm.DB) error {
		var arrAppDishConfirm []ylsdb.Appdishconfirm = make([]ylsdb.Appdishconfirm, 0, 7)
		// 获取餐段类型
		iItemType := getMeal(obRes.Order.Part)
		// 写入长者消费的餐食数据
		for _, v := range obRes.Details {
			// 获取美食数据ID
			pDishInfo, err := DbGetDishInfo(vatools.SInt(v.Pid))
			if err != nil {
				log.Println("同步业务单号", seq, "菜品数据[", v.Pid, "]获取失败:", err.Error())
				seqState = SEQ_NODISH
				return err
			}
			arrAppDishConfirm = append(arrAppDishConfirm, ylsdb.Appdishconfirm{
				Id:             guid.NewString(),
				MemberId:       pMember.Id,
				DishId:         pDishInfo.Id,
				ItemType:       iItemType,
				OperateId:      dbs.Cfg.DishOperateID,
				ConfirmOperate: "cg",
				CreationTime:   time.Time{},
				Count:          1,
				DishName:       pDishInfo.Name,
			})
		}
		if len(arrAppDishConfirm) > 0 {
			// 批量增加
			result := dbs.ObDB.Create(&arrAppDishConfirm)
			if result.Error != nil {
				return result.Error
			}
		}
		// 先写入业务单号数据
		err = DbWriteOrderInfo(&obRes.Order, SEQ_SUCCESS)
		if err != nil {
			return err
		}
		log.Println(seq, "导入成功！")
		return nil
	})
	if err != nil {
		DbWriteOrderInfo(&obRes.Order, seqState)
	}
	ICountSeqDetail += len(obRes.Details)
	ILoopCountSeqDetail += len(obRes.Details)
	ICountSeq++
	ILoopCountSeq++
}

// 将智盘的餐段类型转换为慈爱系统的餐段类型
func getMeal(strMealType string) int {
	switch strMealType {
	case "1":
		return 0
	case "2":
		return 1
	case "3":
		return 2
	default:
		return 100
	}
}
