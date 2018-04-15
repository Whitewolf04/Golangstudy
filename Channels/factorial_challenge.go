package main

import "fmt"

func factorial(n int) chan int {
	total := 1
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}

func main() {
	c := factorial(4)
	for v := range c {
		fmt.Println(v)
	}
}
