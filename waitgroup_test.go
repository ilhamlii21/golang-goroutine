package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchonous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hi")
	time.Sleep(2 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchonous(group)
	}

	group.Wait()
	fmt.Println("Done")
}
