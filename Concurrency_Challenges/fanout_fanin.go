package main

import (
	"fmt"
	"sync"
)

func gen1() chan int {
	v := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 50; j++ {
				v <- j
			}
		}
		close(v)
	}()
	return v
}

func gen2() chan int {
	v := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 50; j++ {
				v <- j
			}
		}
		close(v)
	}()
	return v
}

func withdraw(number chan int) chan int {
	out := make(chan int)
	go func() {
		for m := range number {
			out <- factorial(m)
		}
		close(out)
	}()
	return out
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(channels ...chan int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, n := range channels {
		go func(ch chan int) {
			for m := range ch {
				c <- m
			}
			wg.Done()
		}(n)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func main() {
	f1 := withdraw(gen1())
	f2 := withdraw(gen2())
	for n := range merge(f1, f2) {
		fmt.Println(n)
	}
}
