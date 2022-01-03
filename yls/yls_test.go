package yls

import (
	"github.com/enablefzm/gotools/guid"
	"log"
	"sync"
	"testing"
)

func TestYls(t *testing.T) {
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
}
