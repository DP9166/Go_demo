package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32 = 1
	atomic.AddInt32(&i , 1)
	fmt.Println("i = i + 1 =", i)
	atomic.AddInt32(&i, -1)
	fmt.Println("i = i - 1 =", i)
}