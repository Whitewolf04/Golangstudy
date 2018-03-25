package main

import (
	"fmt"
)

func main() {
	var a, b, c int = 0, 1, 1
	var sum int
	for c < 4000000 {
		if a+b == c {
			a = b
			b = c
			fmt.Println(c)
			if c%2 == 0 {
				sum += c
			}
		}
		c++
	}
	fmt.Println(sum)
}
