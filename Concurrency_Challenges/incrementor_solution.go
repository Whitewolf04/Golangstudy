package main

import (
	"fmt"
	"strconv"
)

func incrementorx(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+ " printing: ", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	return c
}

func main() {
	c := incrementorx(2)
	var countx int

	for n := range c {
		countx++
		fmt.Println(n)
	}

	fmt.Println("Final Counter: ", countx)
}
