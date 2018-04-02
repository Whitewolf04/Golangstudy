package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wwg sync.WaitGroup
var counterx int
var mutex sync.Mutex

func main() {
	wwg.Add(2)
	go incrementore("Foo: ")
	go incrementore("Bar: ")
	wwg.Wait()
	fmt.Println("Final Counter:", counterx)
}

func incrementore(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		counterx++
		fmt.Println(s, i, "Counter:", counterx)
		mutex.Unlock()
	}
	wwg.Done()
}
