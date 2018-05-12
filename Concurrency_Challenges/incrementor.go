package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var wgg sync.WaitGroup

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+" printing:", i)
	}
	wgg.Done()
}

func main() {
	wgg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wgg.Wait()
	fmt.Println("Final Counter: ", count)
}
