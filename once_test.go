package golanggoroutine

import (_
	"fmt"

)

var counter = 0

func OnlyOnce() {
	counter++
}