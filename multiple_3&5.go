package main

import (
	"fmt"
)

func main() {
	var i, product int = 0, 0
	for i < 1000 {
		if i%3 == 0 || i%5 == 0 {
			product = product + i
		}
		i++
	}
	fmt.Println(product)
}
