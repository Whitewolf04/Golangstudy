package main

import "fmt"

func gen(num ...int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range num {
			out <- v
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	c := gen(3, 4, 6)
	cSq := sq(c)
	fmt.Println(<-cSq)
	fmt.Println(<-cSq)
	fmt.Println(<-cSq)

	for n := range sq(sq(gen(5, 7, 28))) {
		fmt.Println(n)
	}
}
