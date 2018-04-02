package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ww sync.WaitGroup
var counter int

func main() {
	ww.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	ww.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	ww.Done()
}
