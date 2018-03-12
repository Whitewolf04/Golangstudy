package main

import "fmt"

func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println("Result: ", counter)
}
