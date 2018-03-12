package main

import "fmt"

func half(a int) (int, bool) {
	return a / 2, bool(a%2 == 0)
}
func main() {
	fmt.Println(half(64))
}
