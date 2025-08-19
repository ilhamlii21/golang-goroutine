package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var Locker = sync.Mutex{}
var cond = sync.NewCond(&Locker)

var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)

	cond.L.Unlock()
	group.Done()
}
func TestCond(t *testing.T) {
	for i := 0; i < 20; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()

		}
	}()

	group.Wait()

}
