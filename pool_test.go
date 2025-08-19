package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Data Baru"
		},
	}

	pool.Put("Ilham")
	pool.Put("Lii")
	pool.Put("Assidaq")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(2 * time.Second)
			pool.Put(data)
		}()

	}
	time.Sleep(12 * time.Second)
	fmt.Println("Selesai mengambil data dari pool")
}
