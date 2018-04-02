package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var gg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	gg.Add(2)
	go fool()
	go bard()
	gg.Wait()
}

func fool() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	gg.Done()
}

func bard() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	gg.Done()
}
