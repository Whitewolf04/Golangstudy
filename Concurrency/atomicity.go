package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wwa sync.WaitGroup
var countera int64

func main() {
	wwa.Add(2)
	go incrementora("Foo: ")
	go incrementora("Bar: ")
	wwa.Wait()
	fmt.Println("Final Counter:", countera)
}

func incrementora(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		//countera++
		atomic.AddInt64(&countera, 1)
		fmt.Println(s, i, "Counter:", countera)
	}
	wwa.Done()
}
