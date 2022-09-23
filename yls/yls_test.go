package yls

import (
	"sync"
	"testing"
)

func TestYls(t *testing.T) {
	/*
		wg := new(sync.WaitGroup)
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				log.Println(guid.NewString())
				guid.NewString()
				wg.Done()
			}()
		}
		wg.Wait()
	*/
	/*
		zpdishseq := ylsdb.Appzpdishseq{
			Id:         guid.NewString(),
			Seq:        "100100101016",
			CreateDate: vatools.STime("2021-10-10 09:00:00"),
			Type:       "ok",
			Amt:        "10",
			Qty:        "20",
			Part:       "2",
			State:      1001,
		}
		result := dbs.ObDB.Create(&zpdishseq)
		log.Println(result.Error)
		log.Println(result.RowsAffected)
	*/
	/*
		var seqs []ylsdb.Appzpdishseq
		// 查询
		dbs.ObDB.Where("seq Like ?", "%10010010101%").Limit(1).Find(&seqs)
		fmt.Println(seqs)
	*/
	/*
		pMember, err := DbGetMember("59210065")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(pMember.FullName)
		}
	*/

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			WorkSyncDish()
			wg.Done()
		}()
	}
	wg.Wait()
	// WorkSyncDish()
}
