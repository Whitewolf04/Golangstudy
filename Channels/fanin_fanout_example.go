package main

import (
	"sync"
	"fmt"
)

func main(){
	c := generate(4, 7)
	c1 := square(c)
	c2 := square(c)

	for n := range merge(c1, c2){
		fmt.Println(n)
	}
}

func generate(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

func square(input chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range input {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func merge(input ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(input))

	for _, c := range input {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
