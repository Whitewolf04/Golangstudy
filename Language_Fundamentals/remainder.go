package main

import "fmt"

func main() {
	x := 547 / 17
	fmt.Println(x)

	var y = &x
	*y = 547 % 17
	fmt.Println(x)
}
